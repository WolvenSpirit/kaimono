import { inherits } from 'util';

export class Create {
    private operation:string = "create";
    public resources:string[];
    public define:string[];
    encode() {
        return JSON.stringify({
            Operation: this.operation,
            Table:this.resources,
            Define:this.define
        });
    }
}
export class Insert {
    private operation = "insert";
    public resources:string[];
    public Cols:string[];
    public Vals:string[];
    encode() {
        return JSON.stringify({
            Operation: this.operation,
            Table:this.resources,
            Insert:{
                Cols:this.Cols,
                Values:this.Vals
            }
        });
    }
}
export class Update {
    private operation = "update";
    public resources:string[];
    public updateAssignments:string[];
    public constraints:string[];
    encode() {
        return JSON.stringify({
            Operation: this.operation,
            Table:this.resources,
            UpdateAssignment:this.updateAssignments,
            Where:this.constraints
        });
    }
}
export class Select {
    private operation = "select";
    public resources:string[];
    public want:string[];
    public constraints:string[];
    encode() {
        return JSON.stringify({
            Operation: this.operation,
            Table:this.resources,
            Want:this.want,
            Where:this.constraints
        });
    }
}
export class Delete {
    private operation = "delete";
    public resources:string[];
    public constraints:string[];
    encode() {
        return JSON.stringify({
            Operation: this.operation,
            Table:this.resources,
            Where:this.constraints
        });
    }
}