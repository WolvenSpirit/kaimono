import { Injectable } from '@angular/core';
import { AppConfig } from './config';
import { HttpClient, HttpHeaders } from '@angular/common/http';
const { HandshakeRequest,HandshakeResponse,ResourceRequest,ResourceOverviewResponse,LoginRequest,AuthResponse,RegistrationRequest,join,insert,RequestBody,ResponseBody } = require("./proto/kaimono_pb");
const { ApiServiceClient } = require("./proto/kaimono_grpc_web_pb");
export { HandshakeRequest,HandshakeResponse,ResourceRequest,ResourceOverviewResponse,LoginRequest,AuthResponse,RegistrationRequest,join,insert,RequestBody,ResponseBody };
@Injectable({
  providedIn: 'root'
})
export class ApiService {
  public kaimonoClient;
  public config:AppConfig;
  constructor(private http:HttpClient) {
    this.config = new AppConfig();
    this.kaimonoClient = new ApiServiceClient(this.config.api_url);
    this.handshake()
  }
   create(request) {
    this.kaimonoClient.create(request,{},(e,r)=>{
      if(!e) {
        return r
      }
    })
   }
   insert(request) {
    this.kaimonoClient.insert(request,{},(e,r)=>{
      if(!e) {
        return r
      }
    })
   }
   update(request) {
    this.kaimonoClient.update(request,{},(e,r)=>{
      if(!e) {
        return r
      }
    })
   }
   select(request) {
    this.kaimonoClient.select(request,{},(e,r)=>{
      if(!e) {
        return r
      }
    })
   }
   delete(request) {
    this.kaimonoClient.delete(request,{},(e,r)=>{
      if(!e) {
        return r
      }
    })
  }
  public claimMode:boolean;
  public handshake() {
    let h = new HandshakeRequest();
    this.kaimonoClient.handshake(h,{},(e,r)=>{
      if(!e) {
        this.claimMode = r.getClaimMode()
        console.log(this.claimMode)
      } else {
        console.log(e)
      }
    });
  }
  public api_resources;
  resources(request) {
    this.kaimonoClient.publicResources(request,{},(e,r)=>{
      if(!e) {
         this.api_resources = r.getResources();
      }else{
        console.log(e);
      }
    })
  }
}