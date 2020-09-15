/*
*  .---------------------------.
* |  Frontend app configuration |
*  '---------------------------'
*/
export class AppConfig {
    public api_protocol:string = "https"
    public api_path:string = "api";
    public api_host:string = "localhost"
    public api_port:string = "8081"
    public api_url:string;
    constructor() {
        this.api_url = `${this.api_protocol}://${this.api_host}:${this.api_port}}`;
    }
}