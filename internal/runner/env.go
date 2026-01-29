package runner

import "os"

func mergeEnv(global, task map[string]string) []string {
	env := make(map[string]string)

	// Base: process env
	for _, e := range os.Environ() {
		k, v := splitEnv(e)
		env[k] = v
	}

	// Global env
	for k, v := range global {
		env[k] = v
	}

	// Task env (highest priority)
	for k, v := range task {
		env[k] = v
	}

	out := make([]string, 0, len(env))
	for k, v := range env {
		out = append(out, k+"="+v)
	}
	return out
}

func splitEnv(s string) (string, string) {
	for i := 0; i < len(s); i++ {
		if s[i] == '=' {
			return s[:i], s[i+1:]
		}
	}
	return s, ""
}
