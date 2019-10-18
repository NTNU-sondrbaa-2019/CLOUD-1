# How to use CO1Cache (with example)

In the main.go file inside /cmd/CO1/ we write diagnostics to the cache.

	diagnostics := CO1Struct.Diagnostics{
		200,
		200,
		"v1",
		0,
	}

	CO1Cache.Initialize()
	CO1Cache.WriteJSON("diagnostics", diagnostics)

**We can add the following after the previous code:**

To properly verify the file exists we start by verifying its existence.

	if CO1Cache.Verify("diagnostics") {
	
Then we defined the output variable and **json.Unmarshal** the diagnostics file into the *referenced* output variable.

		var d2 CO1Struct.Diagnostics
		err := json.Unmarshal(CO1Cache.Read("diagnostics"), &d2)
		
We look for errors and handle any errors that may have occurred.

		if err != nil {
			fmt.Printf("Unable to read diagnostics: %v\n", err)
		}
		
We can now use the output variable as expected.
		
		fmt.Println(d2.Version)
	}

This example without comments:

	if CO1Cache.Verify("diagnostics") {

		var d2 CO1Struct.Diagnostics
		err := json.Unmarshal(CO1Cache.Read("diagnostics"), &d2)

		if err != nil {
			fmt.Printf("Unable to read diagnostics: %v\n", err)
		}
		
		fmt.Println(d2.Version)
	}
