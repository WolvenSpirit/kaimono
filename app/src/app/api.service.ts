import { Injectable } from '@angular/core';
import { AppConfig } from './config';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Create, Insert, Update, Select, Delete } from './payload.frames';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ApiService {
  public config:AppConfig;
  constructor(private http:HttpClient) {
    this.config = new AppConfig();
   }

   post(body) {
     let h = new HttpHeaders();
     h.set("Content-Type","application/json")
    return this.http.post(this.config.api_url,body,{headers:h});
   }

   create(define:string[],resource:string): Observable<any> {
      let c = new Create();
      c.define = define;
      c.resources = [resource];
     return this.post(c.encode());
   }
   insert(cols:string[],vals:string[],resource:string) {
    let i = new Insert();
    i.resources = [resource];
    i.Cols = cols;
    i.Vals = vals;
    return this.post(i.encode());
   }
   update(resource:string,assignments:string[],constraints:string[]) {
    let u = new Update();
    u.resources = [resource];
    u.updateAssignments = assignments;
    u.constraints = constraints;
    return this.post(u.encode());
   }
   select(resources:string[],want:string[],constraints:string[]) {
    let s = new Select();
    s.resources = resources;
    s.want = want;
    s.constraints = constraints;
    return this.post(s.encode());
   }
   delete(resource:string,constraints:string[]) {
    let d = new Delete();
    d.resources = [resource];
    d.constraints = constraints;
    return this.post(d.encode());
   }
  }
