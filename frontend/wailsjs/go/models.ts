export namespace configs {
	
	export class DatabaseConfig {
	    Id: string;
	    User: string;
	    Host: string;
	    Port: string;
	    Password: string;
	    Type: number;
	
	    static createFrom(source: any = {}) {
	        return new DatabaseConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Id = source["Id"];
	        this.User = source["User"];
	        this.Host = source["Host"];
	        this.Port = source["Port"];
	        this.Password = source["Password"];
	        this.Type = source["Type"];
	    }
	}

}

