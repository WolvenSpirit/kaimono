import { Component, OnInit } from '@angular/core';
import { ApiService, HandshakeRequest } from '../api.service';

@Component({
  selector: 'app-mainview',
  templateUrl: './mainview.component.html',
  styleUrls: ['./mainview.component.css']
})
export class MainviewComponent implements OnInit {
  public title = "WolvenSpirit's gRPC example shop" 
  constructor(public api:ApiService) {
  
   }

  private claimMode() {
    var h = new HandshakeRequest();
      this.api.handshake();
  }

  ngOnInit() {
    this.claimMode();
  }

}
