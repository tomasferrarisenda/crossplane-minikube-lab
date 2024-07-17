const express = require("express");
const { Pool } = require("pg");
const app = express();

// Create a PostgreSQL connection pool
const pool = new Pool({
  user: process.env.DB_USER,
  host: process.env.DB_HOST,
  database: process.env.DB_NAME,
  password: process.env.DB_PASSWORD,
  port: process.env.DB_PORT || 5432,
});

// Function to retrieve and increment the visitor count
async function getAndIncrementVisitorCount() {
  const client = await pool.connect();
  try {
    // Start a transaction
    await client.query("BEGIN");

    // Try to update the existing row
    const updateResult = await client.query(
      "UPDATE visitor_count SET count = count + 1 RETURNING count"
    );

    let count;
    if (updateResult.rows.length > 0) {
      // If a row was updated, get the new count
      count = updateResult.rows[0].count;
    } else {
      // If no row was updated, insert a new row
      const insertResult = await client.query(
        "INSERT INTO visitor_count (count) VALUES (1) RETURNING count"
      );
      count = insertResult.rows[0].count;
    }

    // Commit the transaction
    await client.query("COMMIT");

    return count;
  } catch (error) {
    // If an error occurred, roll back the transaction
    await client.query("ROLLBACK");
    console.error("Error retrieving visitor count:", error);
    throw error;
  } finally {
    // Release the client back to the pool
    client.release();
  }
}

// API endpoint to retrieve visitor count
app.get("/", async (req, res) => {
  try {
    const visitorCount = await getAndIncrementVisitorCount();
    res.json({ count: visitorCount });
  } catch (error) {
    res.status(500).json({ error: "Failed to retrieve visitor count" });
  }
});

// Start the server
const port = process.env.PORT || 3000;
app.listen(port, () => {
  console.log(`Server is running on port ${port}`);
});

// Initialize the database table
async function initializeDatabase() {
  const client = await pool.connect();
  try {
    await client.query(`
      CREATE TABLE IF NOT EXISTS visitor_count (
        id SERIAL PRIMARY KEY,
        count INTEGER NOT NULL DEFAULT 0
      )
    `);
    console.log("Database initialized successfully");
  } catch (error) {
    console.error("Error initializing database:", error);
  } finally {
    client.release();
  }
}

initializeDatabase();


// const express = require("express");
// const Redis = require("ioredis");
// const app = express();

// // Create a Redis client
// const redisClient = new Redis({
//   host: process.env.REDIS_HOST, // Replace with the name of the Redis service
//   port: 6379, // Replace with the Redis port if it's different
// });

// // Retrieve the visitor count from Redis
// async function getAndIncrementVisitorCount() {
//   try {
//     const count = await redisClient.incr("visitor_count");
//     return count;
//   } catch (error) {
//     console.error("Error retrieving visitor count:", error);
//     throw error;
//   }
// }

// // API endpoint to retrieve visitor count
// app.get("/", async (req, res) => {
//   try {
//     const visitorCount = await getAndIncrementVisitorCount();
//     res.json({ count: visitorCount });
//   } catch (error) {
//     res.status(500).json({ error: "Failed to retrieve visitor count" });
//   }
// });

// // Start the server
// const port = process.env.PORT || 3000;
// app.listen(port, () => {
//   console.log(`Server is running on port ${port}`);
// });