package testers

import (
	"fmt"

	"github.com/rzjprogramador/AppMotor_GO/entitys/pessoa"
	"github.com/rzjprogramador/AppMotor_GO/externals/libs_rzj"
)

func TesterPessoa() {
	fmt.Println(
		libs_rzj.UseConvertObjectInJson(pessoa.Pessoa1Seed),
		// prototypes - computados
		pessoa.Pessoa1Seed.NomeCompleto(),
		pessoa.Pessoa2Seed.NomeCompleto(),
	)

}
