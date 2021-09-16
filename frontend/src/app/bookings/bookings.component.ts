import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { UtilityService } from '../utility.service';

@Component({
  selector: 'app-bookings',
  templateUrl: './bookings.component.html',
  styleUrls: ['./bookings.component.css']
})
export class BookingsComponent implements OnInit {
  bookings: any
  constructor(private http: HttpClient, public utilityService: UtilityService, private navigateRouter: Router) { }

  ngOnInit(): void {
    let that = this;
    if (this.utilityService.loggedInObj) {
      this.http.get('http://localhost:8080/api/getbookings/' + this.utilityService.loggedInObj.email).subscribe(function (res) {
        that.bookings = res
        console.log("my bookings", res)
      })
    } else {
      this.navigateRouter.navigateByUrl("/")
    }
  }

}
