package utils

//ValidateCommand - to validate commands and look for flags and other information
func ValidateCommand(cType string, args []string) (bool, bool, []string) {
	switch cType {
	case "explore":
		if len(args) <= 2 {
			return true, false, nil
		}
		if args[2] != "-d" && args[2] != "-D" {
			return false, true, args[2:]
		}
		return true, true, args[2:]
	case "visit":
		if len(args) <= 3 {
			return true, false, nil
		}
		if args[3] != "-d" && args[3] != "-D" {
			return false, true, args[2:]
		}
		return true, true, args[2:]
	default:
		return true, true, nil
	}
}
