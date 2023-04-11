package transactionsMaker

func (tm *TransactionMaker) GetMonthReport(month string) (string, error) {
	return tm.repos.GetMonthReport(month)
}
