package utils

// Vérifie si une erreur est survenue. Si c'est le cas, déclenche une panique avec l'erreur.
func CheckErr(err error) {
    if err != nil {
        panic(err)
    }
}