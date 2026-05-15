# 🚗 Garage Management System - Problem Statement

## Overview

A Garage Management System is used to manage customers, their vehicles, mechanics, and service records in a structured way. It allows the garage to track which vehicle was serviced, by which mechanic, and what service was performed along with cost details.

---

## Core Features

1. **Create Customer**
2. **Add Vehicle to Customer**
3. **Create Service Record**
4. **View Customer Vehicles with Services**
5. **View Service Records with Details**

---

## API Endpoints

### Customer APIs (2 Marks)

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/customers` | Create a new customer |
| GET | `/customers/{id}/vehicles` | Get customer vehicles with services |

---

### Vehicle APIs (1.5 Marks)

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/vehicles` | Add vehicle to a customer |

---

### Service APIs (4.5 Marks)

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/service-records` | Create a service record |
| GET | `/service-records` | Get all service records with details |

---

## Database Schema

### TABLE: Customers

| Column | Type | Constraints |
|--------|------|-------------|
| `id` | Integer | Primary Key |
| `name` | String | - |
| `phone` | String | Unique |
| `email` | String | Unique |
| `created_at` | Timestamp | - |

---

### TABLE: Vehicles

| Column | Type | Constraints |
|--------|------|-------------|
| `id` | Integer | Primary Key |
| `customer_id` | Integer | Foreign Key → Customers.id |
| `number_plate` | String | Unique |
| `model` | String | - |
| `created_at` | Timestamp | - |

---

### TABLE: Mechanics

| Column | Type | Constraints |
|--------|------|-------------|
| `id` | Integer | Primary Key |
| `name` | String | - |

---

### TABLE: Service_Records

| Column | Type | Constraints |
|--------|------|-------------|
| `id` | Integer | Primary Key |
| `vehicle_id` | Integer | Foreign Key → Vehicles.id |
| `mechanic_id` | Integer | Foreign Key → Mechanics.id |
| `service` | String | Required |
| `price` | Float | - |
| `date` | Timestamp | - |

---

## JOIN Requirements (Important)

- Customer → Vehicle  
- Vehicle → ServiceRecord  
- ServiceRecord → Mechanic  

All GET APIs must use **JOIN or GORM Preload**.

---

## Edge Cases (2 Marks)

### 🧍 Customer

- Creating customer with duplicate phone/email
- Missing required fields (name, phone)
- Invalid email format

---

### 🚗 Vehicle

- Adding vehicle with invalid customer_id
- Duplicate number_plate
- Missing required fields

---

### 🔧 Service Record

- Creating record with invalid vehicle_id
- Creating record with invalid mechanic_id
- Missing service name
- Negative or zero price

---

### 📋 Fetch APIs

- Invalid ID in path params
- No data found
- Improper JOIN resulting in missing data

---

## ⭐ Features (1.5 Marks)

1. Add filtering in `/service-records` (by mechanic or date)
2. Add pagination (limit, offset)
3. Add sorting (price, date)

---

## Scoring Breakdown

- **Customer APIs**: 2 Marks  
- **Vehicle APIs**: 1.5 Marks  
- **Service APIs**: 4.5 Marks  
- **Edge Cases**: 2 Marks  
- **Bonus Features**: 1.5 Marks  

**Total**: 10 Marks  

---

## Business Rules

### Customer Rules

- Phone must be unique  
- Email must be unique  

---

### Vehicle Rules

- Each vehicle must belong to one customer  
- Number plate must be unique  

---

### Service Record Rules

- Each service must be linked to:
  - One vehicle  
  - One mechanic  

- **Service**: Required field  
- **Price**: Must be greater than 0  
- **Date**: Default = current timestamp  

---

### Data Fetching Rules

- Customer APIs should return only their vehicles  
- Service records must include:
  - Vehicle number plate  
  - Mechanic name  

---

## 🎯 Learning Outcomes

- GORM relationships (One-to-Many)
- JOIN queries in PostgreSQL
- API design using Gin
- Clean schema design
- Backend structuring basics

---