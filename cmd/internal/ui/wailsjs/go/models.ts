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

}

