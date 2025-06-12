package models

import (
	"coinpe/pkg/logger"
	"coinpe/pkg/utils"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	EntityWallet = "wa_"
	EntityINR    = "INR"
)

type Wallet struct {
	ID        uint64         `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	UserUUID            string `json:"user_uuid" gorm:"not null;uniqueIndex"`
	UUID                string `json:"uuid" gorm:"unique;not null"`
	TotalBalanceInCents int    `json:"total_balance_in_cents" gorm:"default:0;not null"`
	Currency            string `json:"currency" gorm:"not null"`

	AdditionalInfo        datatypes.JSON `json:"additional_info"`
	OverdraftLimitInCents uint           `json:"overdraft_limit_in_cents"`
}

type walletRepo struct {
	db *gorm.DB
}

var CoinpeWallet = Wallet{
	ID:                    1,
	UserUUID:              "cpe_XazgAniKetfjJwzkMuMi",
	UUID:                  "wa_vSrAnivBdiKTTpwi8uWx",
	TotalBalanceInCents:   100000,
	Currency:              EntityINR,
	OverdraftLimitInCents: 10000,
}

func (w *Wallet) BeforeCreate(tx *gorm.DB) (err error) {
	tx.Statement.AddClause(clause.OnConflict{
		DoNothing: true,
	})
	if w.UUID == "" {
		w.UUID, err = utils.GenerateNanoID(20, EntityWallet)
		if err != nil {
			return err
		}
	}
	return
}

func (r *walletRepo) Get(where *Wallet) (*Wallet, error) {
	return r.GetWithTx(r.db, where)
}

func (r *walletRepo) GetWithTx(tx *gorm.DB, where *Wallet) (*Wallet, error) {
	var o Wallet
	err := tx.Model(&Wallet{}).Where(where).First(&o).Error
	if err != nil {
		return nil, err
	}
	return &o, nil
}

func (r *walletRepo) Create(u *Wallet) error {
	return r.CreateWithTx(r.db, u)
}

func (r *walletRepo) CreateWithTx(tx *gorm.DB, u *Wallet) error {
	err := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(u).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *walletRepo) Delete(walletID string) error {
	err := r.db.Where(&Wallet{UUID: walletID}).Delete(&Wallet{}).Error
	if err != nil {
		return err
	}
	return nil
}

// func (r *walletRepo) Credit(wallet *Wallet, transaction *Transaction, purposeCode purposecodes.TransactionPurposeCode) (*Wallet, error) {
// 	tx := r.db.Begin()
// 	w, err := r.CreditWithTx(tx, wallet, transaction, purposeCode)
// 	if err != nil {
// 		logger.Error(err)
// 		return wallet, err
// 	}

// 	err = tx.Commit().Error
// 	if err != nil {
// 		logger.Error(err)
// 		return w, err
// 	}

// 	return w, nil
// }

// func (r *walletRepo) CreditWithTx(tx *gorm.DB, wallet *Wallet, transaction *Transaction, purposeCode purposecodes.TransactionPurposeCode) (*Wallet, error) {
// 	transactionRepo := InitTransactionRepo(r.db)
// 	updatedWalletBalance := wallet.TotalBalanceInCents + transaction.AmountInCents
// 	transaction.Type = constants.TransactionTypeCredit
// 	transaction.WalletID = wallet.ID
// 	transaction.OpeningBalanceInCents = wallet.TotalBalanceInCents
// 	transaction.ClosingBalanceInCents = updatedWalletBalance
// 	transaction.Status = constants.EntitySuccess
// 	transaction.ToWalletUUID = wallet.UUID
// 	transaction.PurposeCode = purposeCode

// 	err := transactionRepo.CreateWithTx(tx, transaction)
// 	if err != nil {
// 		tx.Rollback()
// 		logger.Error(err)
// 		return wallet, err
// 	}

// 	err = tx.Model(&wallet).Select("total_balance_in_cents").Updates(map[string]interface{}{"total_balance_in_cents": updatedWalletBalance}).Error
// 	if err != nil {
// 		tx.Rollback()
// 		logger.Error(err)
// 		return wallet, err
// 	}

// 	wallet.TotalBalanceInCents = updatedWalletBalance
// 	return wallet, nil
// }

// func (r *walletRepo) Debit(wallet *Wallet, transaction *Transaction, purposeCode purposecodes.TransactionPurposeCode) (*Wallet, error) {
// 	tx := r.db.Begin()
// 	w, err := r.DebitWithTx(tx, wallet, transaction, purposeCode)
// 	if err != nil {
// 		logger.Error(err)
// 		return wallet, err
// 	}

// 	err = tx.Commit().Error
// 	if err != nil {
// 		logger.Error(err)
// 		return w, err
// 	}

// 	return w, nil
// }

// func (r *walletRepo) DebitWithTx(tx *gorm.DB, wallet *Wallet, transaction *Transaction, purposeCode purposecodes.TransactionPurposeCode) (*Wallet, error) {
// 	if !r.CanDebit(wallet, transaction.AmountInCents) {
// 		return nil, errors.New("insufficient funds")
// 	}

// 	transactionRepo := InitTransactionRepo(r.db)
// 	updatedWalletBalance := wallet.TotalBalanceInCents - transaction.AmountInCents
// 	transaction.Type = constants.TransactionTypeDebit
// 	transaction.WalletID = wallet.ID
// 	transaction.OpeningBalanceInCents = wallet.TotalBalanceInCents
// 	transaction.ClosingBalanceInCents = updatedWalletBalance
// 	transaction.Status = constants.EntitySuccess
// 	transaction.FromWalletUUID = wallet.UUID
// 	transaction.PurposeCode = purposeCode

// 	err := transactionRepo.CreateWithTx(tx, transaction)
// 	if err != nil {
// 		tx.Rollback()
// 		logger.Error(err)
// 		return wallet, err
// 	}

// 	err = tx.Model(&wallet).Updates(map[string]interface{}{"total_balance_in_cents": updatedWalletBalance}).Error
// 	if err != nil {
// 		tx.Rollback()
// 		logger.Error(err)
// 		return wallet, err
// 	}

// 	wallet.TotalBalanceInCents = updatedWalletBalance
// 	return wallet, nil
// }

func (r *walletRepo) UpdateBalance(tx *gorm.DB, walletID string, amountInCents int) error {
	err := tx.Model(&Wallet{}).Where(&Wallet{UUID: walletID}).Updates(map[string]interface{}{"total_balance_in_cents": amountInCents}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *walletRepo) CanDebit(wallet *Wallet, amountInCents int) bool {
	return amountInCents <= (wallet.TotalBalanceInCents + int(wallet.OverdraftLimitInCents))
}

// Update implements IWallet.
func (r *walletRepo) Update(where *Wallet, w *Wallet) error {
	return r.UpdateWithTx(r.db, where, w)
}

// UpdateWithTx implements IWallet.
func (r *walletRepo) UpdateWithTx(tx *gorm.DB, where *Wallet, w *Wallet) error {
	err := tx.Model(&Wallet{}).Where(where).Updates(w).Error
	if err != nil {
		logger.Error("unable to update wallet ", err)
		return err
	}
	return nil
}
