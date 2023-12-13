Com &cobra.Command{Use: "go-cli-cobra-sample"} ao instalar o sh, é possível chamar o sh pelo comando "go-cli-cobra-sample" no terminal

Cada cmd é um comando que é adicionado ao CLI, Ex: Create, Delete, Etc...

O `cmd.Flags().StringVarP()` a realizar o 'de-para' entre os argumentos do comando no acionamento e as variáveis que são utilizadas dentro da lógica, usando ponteiros.

```console
./cobra-pkg-sample create -e meireles.luhan@gmail.com -n Luhan -s 123456
Nome: Luhan
Email: meireles.luhan@gmail.com
Senha: 123456
```

Referência: https://www.youtube.com/watch?v=Ph2QEB0WMcA
