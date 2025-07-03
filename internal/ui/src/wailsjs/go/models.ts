export namespace data {
	
	export class Expense {
	    ID: number;
	    ProjectID: number;
	    Category: string;
	    Amount: number;
	
	    static createFrom(source: any = {}) {
	        return new Expense(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.ProjectID = source["ProjectID"];
	        this.Category = source["Category"];
	        this.Amount = source["Amount"];
	    }
	}
	export class Income {
	    ID: number;
	    ProjectID: number;
	    Source: string;
	    Amount: number;
	
	    static createFrom(source: any = {}) {
	        return new Income(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.ProjectID = source["ProjectID"];
	        this.Source = source["Source"];
	        this.Amount = source["Amount"];
	    }
	}
	export class Project {
	    ID: number;
	    Name: string;
	
	    static createFrom(source: any = {}) {
	        return new Project(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Name = source["Name"];
	    }
	}

}

