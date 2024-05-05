# Whoshittin: NYC Jazz Clubs

## Overview

Whoshittin is a data aggregator designed to gather information about upcoming jazz performances at various clubs in New York City. It extracts data such as artist names, performance times, and band members and stores it in a MongoDB database. Additionally, it provides an API for interacting with the scraped data, allowing users to retrieve performance information.

## Features

- **Aggregator**: Automatically collects data from New York City jazz clubs to gather information about upcoming performances.
- **API**: 
  - **Read-Only Endpoints**: Provides endpoints for users to retrieve performance data. These endpoints are read-only and intended for external users.
  - **Writeable Endpoints**: Provides endpoints for internal use to write performance data to the database. These endpoints are not exposed to external users.
- **Data Enrichment**: Extracts key details such as artist names, performance times, and band members from scraped web pages.
- **Cron Job Integration**: Data aggregator runs on a scheduled interval using a cron job to ensure up-to-date performance information.
- **Website**: Renders the performance data in a user interface, fetching the data from the Whoshittin API.

