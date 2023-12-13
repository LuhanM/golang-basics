package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	var roodCmd = &cobra.Command{Use: "go-cli-cobra-sample"}

	var nome, email, senha string

	var cmd = &cobra.Command{
		Use:   "create",
		Short: "Crie um novo usuário",
		Run: func(cmd *cobra.Command, args []string) {
			if nome == "" {
				fmt.Println("nome não pode estar vazio.")
				return
			}

			if email == "" {
				fmt.Println("email não pode estar vazio.")
				return
			}

			if len(senha) < 6 {
				fmt.Println("A senha deve possui ao menos 6 caracteres.")
				return
			}

			fmt.Printf("Nome: %s\nEmail: %s\nSenha: %s\n", nome, email, senha)
		},
	}

	cmd.Flags().StringVarP(&nome, "nome", "n", "", "Nome do usuário")
	cmd.Flags().StringVarP(&email, "email", "e", "", "Email do usuário")
	cmd.Flags().StringVarP(&senha, "senha", "s", "", "Senha do usuário")

	roodCmd.AddCommand(cmd)
	roodCmd.Execute()
}
