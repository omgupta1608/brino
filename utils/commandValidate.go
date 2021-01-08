package utils

//ValidateCommand - to validate commands and look for flags and other information
func ValidateCommand(cType string, args []string) (bool, bool, []string) {
	if len(args) <= 2 {
		return true, false, nil
	}
	switch cType {
	case "explore", "visit":
		if args[2] != "-d" && args[2] != "-D" {
			return false, true, args[2:]
		}
		return true, true, args[2:]
	default:
		return true, true, nil
	}
}
