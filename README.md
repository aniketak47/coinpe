# CoinPe

**CoinPe** is a digital wallet ledger system that enables users to manage transactions using virtual coins. It provides a robust and scalable solution for tracking coin-based credits and debits, making it ideal for loyalty programs, internal credit systems, gamified platforms, or custom virtual economies.

---

## 🚀 Features

- 🔐 User authentication and account management  
- 💰 Create coin-based credit and debit transactions  
- 📊 Real-time wallet balance tracking  
- 🧾 View detailed transaction history  
- 🏷️ Add metadata to transactions for better organization  
- 🔁 Supports multiple transaction types (credit, debit, reward, etc.)  
- 🔌 API-first design for easy integration with other systems

---

## 🧑‍💻 Tech Stack

- **Backend:** Golang (Gin)  
- **Database:** PostgreSQL
- **Authentication:** Supabase / JWT / Custom Auth  
- **Deployment:** Docker / AWS 

---

## 📦 Installation

### Prerequisites
- Go >= 1.20  
- PostgreSQL or any SQL-compatible DB  
- Git  
- (Optional) Docker

### Steps

```bash
# Clone the repository
git clone https://github.com/aniketak47/coinpe.git
cd coinpe

# Install dependencies
go mod tidy

# Set environment variables
cp .env.example .env
# Fill in the required environment variables

# Run the application
go run main.go
