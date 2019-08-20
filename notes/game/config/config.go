package config

ipmort "zpatricks/foothing/config"

// todo: Register(context?) error

func Register() error {
	schema := config.ConfigSchema{
		Name: "prometheus",
		Inputs: []config.InputSchema{
			&config.IntInputSchema{
				Name: "port",
				Default: 9090,
				Description: "The port to listen on.",
				Readers: []config.InputReader{
					config.EnvironmentVariableReader{
						Name: "PROM_PORT"
					},
					config.CLIFlagReader{
						Name: "port",
						Aliases: []string{"p"},
					},
				},
				// Always read via EnvVars and/or CLI Flags. Values are populated via Environment (NoOp) and os.Args. 
				// When in ci/cd, get it using a key in Jenkins.    
				// When in staging, get it using a load balancer url from k8s.
				// Do the same in prod, but from different cluster. 
			},
		},	
	}
	
	return config.Register()
}

var defaultSchema = config.Schema {
	Version: version,
	Inputs: []config.Input{
		// prod vs. dev contexts?
		config.IntInput{
			Name: "port",
			Default: 9090,
			Description: "The port to listen on.",
			Generators: []config.IntGenerator{
				EnvironmentVariableGenerator{
					EnvironmentVariable: "APP_PORT",
				},
				CLIFlagGeneratorInt{
					FlagName: "port",
					Aliases: []string{"p"},		
					// Need another wire-up to ensure CLIFlagGenerator gets parsed w/ the os.Args. 
					// Will want to create own CLI library for this, but I think it'll be worth it in the end. 
					// keep it lightweight, only what I need - e.g. Name, Default, Description. 
				},
			},
		},
	},
}


/*
	If you tihnk about it, most apps are just flipping switches somewhere on a backend.	
	That's all user input is: textbox, save input as each letter types.  
		
*/
