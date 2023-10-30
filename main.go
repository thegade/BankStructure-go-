package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type BankAccount struct {
	AccountNumber string  `json:"account_number"`
	Balance       float64 `json:"balance"`
	Currency      string  `json:"currency"`
}

type Card struct {
	CardNumber string `json:"card_number"`
	Account    string `json:"account"`
}

type Transaction struct {
	FromAccount string  `json:"from_account"`
	ToAccount   string  `json:"to_account"`
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
}

func main() {
	bankAccounts := []BankAccount{
		{AccountNumber: "11111111", Balance: 1000, Currency: "USD"},
		{AccountNumber: "22222222", Balance: 500, Currency: "EUR"},
	}

	cards := []Card{
		{CardNumber: "1111-1111-1111-1111", Account: "11111111"},
		{CardNumber: "2222-2222-2222-2222", Account: "22222222"},
	}

	transactions := []Transaction{}

	for {
		fmt.Println("Выберите пункт действия:")
		fmt.Println("1) Управление банковскими счетами")
		fmt.Println("2) Управление картами")
		fmt.Println("3) Денежные операции")
		fmt.Println("4) Валютные операции")
		fmt.Println("5) Выйти")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			bankAccountManagement(&bankAccounts)
		case 2:
			cardManagement(&cards, &bankAccounts)
		case 3:
			moneyOperations(&bankAccounts, &transactions)
		case 4:
			currencyOperations(&bankAccounts, &transactions)
		case 5:
			saveData(bankAccounts, cards, transactions)
			return
		default:
			fmt.Println("Некорректный выбор")
		}
	}
}

func bankAccountManagement(bankAccounts *[]BankAccount) {
	fmt.Println("Выберите пункт действия:")
	fmt.Println("1) Просмотреть баланс")
	fmt.Println("2) Создать новый счет")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		fmt.Println("Введите номер счета:")
		var accountNumber string
		fmt.Scanln(&accountNumber)

		for _, account := range *bankAccounts {
			if account.AccountNumber == accountNumber {
				fmt.Printf("Баланс на счете %s: %.2f %s\n", accountNumber, account.Balance, account.Currency)
				return
			}
		}
		fmt.Println("Счет с таким номером не найден")

	case 2:
		fmt.Println("Введите номер нового счета:")
		var accountNumber string
		fmt.Scanln(&accountNumber)

		for _, account := range *bankAccounts {
			if account.AccountNumber == accountNumber {
				fmt.Println("Счет с таким номером уже существует")
				return
			}
		}
		fmt.Println("Введите начальный баланс:")
		var balance float64
		fmt.Scanln(&balance)

		fmt.Println("Введите валюту:")
		var currency string
		fmt.Scanln(&currency)

		*bankAccounts = append(*bankAccounts, BankAccount{AccountNumber: accountNumber, Balance: balance, Currency: currency})
		fmt.Println("Новый счет успешно создан")
	default:
		fmt.Println("Некорректный выбор")
	}
}

func cardManagement(cards *[]Card, bankAccounts *[]BankAccount) {
	fmt.Println("Выберите пункт действия:")
	fmt.Println("1) Просмотреть список карт")
	fmt.Println("2) Выдать новую карту")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		for _, card := range *cards {
			fmt.Printf("Карта: %s, Привязано к счету: %s\n", card.CardNumber, card.Account)
		}
	case 2:
		fmt.Println("Введите номер счета:")
		var accountNumber string
		fmt.Scanln(&accountNumber)

		var accountExists bool
		for _, account := range *bankAccounts {
			if account.AccountNumber == accountNumber {
				accountExists = true
				break
			}
		}
		if !accountExists {
			fmt.Println("Счет с таким номером не найден")
			return
		}

		fmt.Println("Введите номер новой карты:")
		var cardNumber string
		fmt.Scanln(&cardNumber)

		for _, card := range *cards {
			if card.CardNumber == cardNumber {
				fmt.Println("Карта с таким номером уже существует")
				return
			}
		}

		*cards = append(*cards, Card{CardNumber: cardNumber, Account: accountNumber})
		fmt.Println("Новая карта успешно выдана")
	default:
		fmt.Println("Некорректный выбор")
	}
}

func moneyOperations(bankAccounts *[]BankAccount, transactions *[]Transaction) {
	fmt.Println("Выберите пункт действия:")
	fmt.Println("1) Пополнить счет")
	fmt.Println("2) Снять деньги со счета")
	fmt.Println("3) Перевести деньги между счетами")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		fmt.Println("Введите номер счета:")
		var accountNumber string
		fmt.Scanln(&accountNumber)

		for i := range *bankAccounts {
			if (*bankAccounts)[i].AccountNumber == accountNumber {
				fmt.Println("Введите сумму для пополнения:")
				var amount float64
				fmt.Scanln(&amount)

				(*bankAccounts)[i].Balance += amount
				fmt.Printf("Счет %s успешно пополнен на %.2f %s\n", accountNumber, amount, (*bankAccounts)[i].Currency)
				return
			}
		}
		fmt.Println("Счет с таким номером не найден")

	case 2:
		fmt.Println("Введите номер счета:")
		var accountNumber string
		fmt.Scanln(&accountNumber)

		for i := range *bankAccounts {
			if (*bankAccounts)[i].AccountNumber == accountNumber {
				fmt.Println("Введите сумму для снятия:")
				var amount float64
				fmt.Scanln(&amount)

				if amount <= (*bankAccounts)[i].Balance {
					(*bankAccounts)[i].Balance -= amount
					fmt.Printf("Со счета %s успешно снято %.2f %s\n", accountNumber, amount, (*bankAccounts)[i].Currency)
				} else {
					fmt.Println("Недостаточно средств на счете")
				}
				return
			}
		}
		fmt.Println("Счет с таким номером не найден")

	case 3:
		fmt.Println("Введите номер счета отправителя:")
		var fromAccount string
		fmt.Scanln(&fromAccount)

		fmt.Println("Введите номер счета получателя:")
		var toAccount string
		fmt.Scanln(&toAccount)

		var senderExists bool
		var receiverExists bool
		for i := range *bankAccounts {
			if (*bankAccounts)[i].AccountNumber == fromAccount {
				senderExists = true
			}
			if (*bankAccounts)[i].AccountNumber == toAccount {
				receiverExists = true
			}
		}
		if !senderExists {
			fmt.Println("Счет отправителя не найден")
			return
		}
		if !receiverExists {
			fmt.Println("Счет получателя не найден")
			return
		}

		fmt.Println("Введите сумму перевода:")
		var amount float64
		fmt.Scanln(&amount)

		var senderCurrency string
		var receiverCurrency string

		for _, account := range *bankAccounts {
			if account.AccountNumber == fromAccount {
				senderCurrency = account.Currency
			}
			if account.AccountNumber == toAccount {
				receiverCurrency = account.Currency
			}
		}

		if senderCurrency != receiverCurrency {
			fmt.Println("Перевод между различными валютами недопустим")
			return
		}

		if amount <= 0 {
			fmt.Println("Сумма перевода должна быть положительной")
			return
		}

		for i := range *bankAccounts {
			if (*bankAccounts)[i].AccountNumber == fromAccount {
				if amount > (*bankAccounts)[i].Balance {
					fmt.Println("Недостаточно средств на счете отправителя")
					return
				}
				(*bankAccounts)[i].Balance -= amount
			}
			if (*bankAccounts)[i].AccountNumber == toAccount {
				(*bankAccounts)[i].Balance += amount
			}
		}

		*transactions = append(*transactions, Transaction{FromAccount: fromAccount, ToAccount: toAccount, Amount: amount, Currency: senderCurrency})
		fmt.Printf("Перевод успешно выполнен: %s -> %s, %.2f %s\n", fromAccount, toAccount, amount, senderCurrency)

	default:
		fmt.Println("Некорректный выбор")
	}
}

func currencyOperations(bankAccounts *[]BankAccount, transactions *[]Transaction) {
	fmt.Println("Выберите пункт действия:")
	fmt.Println("1) Перевести деньги между счетами в различных валютах")
	fmt.Println("2) Конвертировать сумму из одной валюты в другую")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		fmt.Println("Введите номер счета отправителя:")
		var fromAccount string
		fmt.Scanln(&fromAccount)

		fmt.Println("Введите номер счета получателя:")
		var toAccount string
		fmt.Scanln(&toAccount)

		fmt.Println("Введите сумму перевода:")
		var amount float64
		fmt.Scanln(&amount)

		var senderCurrency string
		var receiverCurrency string

		for _, account := range *bankAccounts {
			if account.AccountNumber == fromAccount {
				senderCurrency = account.Currency
			}
			if account.AccountNumber == toAccount {
				receiverCurrency = account.Currency
			}
		}

		if senderCurrency == receiverCurrency {
			fmt.Println("Перевод между одинаковыми валютами недопустим")
			return
		}

		if amount <= 0 {
			fmt.Println("Сумма перевода должна быть положительной")
			return
		}

		var exchangeRate float64

		switch senderCurrency {
		case "USD":
			switch receiverCurrency {
			case "EUR":
				exchangeRate = 0.85
			default:
				fmt.Println("Конвертация между указанными валютами невозможна")
				return
			}
		case "EUR":
			switch receiverCurrency {
			case "USD":
				exchangeRate = 1.18
			default:
				fmt.Println("Конвертация между указанными валютами невозможна")
				return
			}
		default:
			fmt.Println("Конвертация между указанными валютами невозможна")
			return
		}

		convertedAmount := amount * exchangeRate

		for i := range *bankAccounts {
			if (*bankAccounts)[i].AccountNumber == fromAccount {
				if amount > (*bankAccounts)[i].Balance {
					fmt.Println("Недостаточно средств на счете отправителя")
					return
				}
				(*bankAccounts)[i].Balance -= amount
			}
			if (*bankAccounts)[i].AccountNumber == toAccount {
				(*bankAccounts)[i].Balance += convertedAmount
			}
		}

		*transactions = append(*transactions, Transaction{FromAccount: fromAccount, ToAccount: toAccount, Amount: amount, Currency: senderCurrency})
		fmt.Printf("Перевод успешно выполнен: %s -> %s, %.2f %s (%.2f %s)\n", fromAccount, toAccount, amount, senderCurrency, convertedAmount, receiverCurrency)

		fmt.Printf("Перевод успешно выполнен: %s -> %s, %.2f %s (%.2f %s)\n", fromAccount, toAccount, amount, senderCurrency, convertedAmount, receiverCurrency)

	case 2:
		fmt.Println("Введите номер счета:")
		var accountNumber string
		fmt.Scanln(&accountNumber)

		var accountExists bool
		for _, account := range *bankAccounts {
			if account.AccountNumber == accountNumber {
				accountExists = true
				break
			}
		}
		if !accountExists {
			fmt.Println("Счет с таким номером не найден")
			return
		}

		fmt.Println("Введите текущую валюту счета:")
		var currentCurrency string
		fmt.Scanln(&currentCurrency)

		for i, account := range *bankAccounts {
			if account.AccountNumber == accountNumber && account.Currency == currentCurrency {
				fmt.Println("Введите сумму для конвертации:")
				var amount float64
				fmt.Scanln(&amount)

				fmt.Println("Введите целевую валюту:")
				var targetCurrency string
				fmt.Scanln(&targetCurrency)

				var exchangeRate float64

				switch currentCurrency {
				case "USD":
					switch targetCurrency {
					case "EUR":
						exchangeRate = 0.85
					default:
						fmt.Println("Конвертация между указанными валютами невозможна")
						return
					}
				case "EUR":
					switch targetCurrency {
					case "USD":
						exchangeRate = 1.18
					default:
						fmt.Println("Конвертация между указанными валютами невозможна")
						return
					}
				default:
					fmt.Println("Конвертация между указанными валютами невозможна")
					return
				}

				convertedAmount := amount * exchangeRate

				if amount > (*bankAccounts)[i].Balance {
					fmt.Println("Недостаточно средств на счете")
					return
				}

				account.Balance -= amount
				fmt.Printf("Со счета %s успешно конвертировано %.2f %s в %.2f %s\n", accountNumber, amount, currentCurrency, convertedAmount, targetCurrency)
				return
			}
		}
		fmt.Println("Счет с таким номером и указанной валютой не найден")

	default:
		fmt.Println("Некорректный выбор")
	}
}

func saveData(bankAccounts []BankAccount, cards []Card, transactions []Transaction) {
	data := struct {
		BankAccounts []BankAccount `json:"bank_accounts"`
		Cards        []Card        `json:"cards"`
		Transactions []Transaction `json:"transactions"`
	}{BankAccounts: bankAccounts, Cards: cards, Transactions: transactions}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Ошибка при сериализации данных:", err)
		return
	}

	err = os.WriteFile("data.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Ошибка при записи данных в файл:", err)
		return
	}

	fmt.Println("Данные успешно сохранены в файле data.json")
}
