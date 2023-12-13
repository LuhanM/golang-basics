**English**

Fuzz tests should be created within the same XYZ_TEST file.

You Must start with the prefix "Fuzz" example "Fuzzmyindexany" and with the parameters `f *testing.f` Unlike unit tests using` t *testing.t`

To execute the Fuzz Test:
    
    go test -v -fuzz FuzzMyIndexAny

If some test scenario created by fuzz break the tests, this scenario will be saved in "testdata/Fuzz/$fuzz_test_function_name" 
ex. ```fuzz-test/testdata/fuzz/FuzzMyIndexAny/b01eab59595e769c```
___________

**Português**

Testes fuzz devem ser criados dentro do mesmo arquivo xyz_test.

Deve começar com o prefixo "Fuzz" exemplo "FuzzMyIndexAny" e com os parâmetros `f *testing.F` diferentemente dos testes unitários que utilizam `t *testing.T`

Para executar o fuzz test: 

    go test -v -fuzz FuzzMyIndexAny

Se algum cenário de teste criado por Fuzz quebrar os testes, esse cenário será salvo em "testData/fuzz/$ fuzz_test_function_name"