package main

import "fmt"

type SavingAcc struct {
	accNumber int
	holderName string
	balance  int
	interest int
}
type CurrentAcc struct {
	accNumber int
	holderName string
	balance  int
}
type BankOperation interface {
	deposit(int) int
	withdraw(int) int
	checkBalance() int
}
//interestCredited() int
//  Saving Account
func (ac SavingAcc) deposit(am int) int {
	ac.balance += am
	return ac.balance
}
func (ac SavingAcc)withdraw(am int) int {
	if ac.balance < am{
		fmt.Printf("Saving Account balance(%d) is less then requested amount (%d)",ac.balance,am)
		return 0
	}
	ac.balance -= am
	return ac.balance
}
func (ac SavingAcc) checkBalance() int {
	//ac.balance += am
	return ac.balance
}
//  Current Account
func (ac CurrentAcc) deposit(am int) int {
	ac.balance += am
	return ac.balance
}
func (ac CurrentAcc)withdraw(am int) int {
	if ac.balance < am{
		fmt.Printf("Current Account balance(%d) is less then requested amount (%d)",ac.balance,am)
		return 0
	}
	ac.balance -= am
	return ac.balance
}
func (ac CurrentAcc) checkBalance() int {
	//ac.balance += am
	return ac.balance
}
func OpenAccounts() []BankOperation {
	a := make([]BankOperation, 0)
	a1 := SavingAcc{1, "Raju", 100, 0}
	a2 := CurrentAcc{2, "Shaikh", 100}
	a3 := SavingAcc{3, "Amit", 100, 0}
	a4 := SavingAcc{4, "Nilesh", 100, 200}
	a5 := CurrentAcc{5, "Harmain", 500}
	//for i:=range []Account(a1,a2,a3,a4,a5){
	//}
	a = append(a, a1,a2,a3,a4,a5)
	return a

}
func getInterestAmount(acc BankOperation) int {
	v, ok := acc.(SavingAcc)
	if !ok{
		fmt.Println("Only Savings account can have intereset")
		return -1
	}
	fmt.Println("Acc received T : %T", acc)
	//fmt.Println("Acc received v : %v",acc)
	return v.interest
}
func PrintAccDetails(acc BankOperation)  {
	v, ok := acc.(SavingAcc)
	if !ok{
		vc := acc.(CurrentAcc)
		fmt.Printf("Current AccountNo: %d, Balance is : %d \n",vc.accNumber, vc.balance)
		return
	}
	fmt.Printf("Saving AccountNo: %d, Balance is : %d and Interest is : %d\n",v.accNumber, v.balance, v.interest)
}
func main()  {
	ac := OpenAccounts()
	fmt.Println(ac)
	for _, acc :=range ac{
		PrintAccDetails(acc)
	}
	ret := getInterestAmount(ac[3])
	fmt.Printf("Interest received for Acc : %v is %d \n",ac[3], ret)
	ret1 := getInterestAmount(ac[1])
	if ret1==-1{
		fmt.Println("Ohh it is Current account, so don't have interest")
	}
}