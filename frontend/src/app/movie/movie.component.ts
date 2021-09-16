import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { MatSnackBar } from '@angular/material/snack-bar';
import { ActivatedRoute, Router } from '@angular/router';
import * as _ from 'lodash';
import { SnackbarComponent } from '../snackbar/snackbar.component';
import { UtilityService } from '../utility.service';

@Component({
  selector: 'app-movie',
  templateUrl: './movie.component.html',
  styleUrls: ['./movie.component.css']
})
export class MovieComponent implements OnInit {
  movieName: any
  movie: any
  show: any
  movieShows: any = []
  showDetails: any = {}
  isShowTickets: boolean = false
  tickets: any = []
  selectedShow: any
  selectedSeats: any
  spinner: boolean = false
  selectedIndex: any
  totalAmount: any = 0
  seatsCount: any = 0
  userObj: any = {
    email: '',
    movie_name: '',
    date: '',
    movie_time: '',
    seats: []
  }

  couponcode: any
  timeDifference: any
  coupons: any
  couponPrice: any = 0;

  constructor(public utilityService: UtilityService, private activatedRoute: ActivatedRoute, private http: HttpClient,
    private snackbar: MatSnackBar, private navigateRouter: Router) {
    this.movieName = this.activatedRoute.snapshot.paramMap.get('movie');
  }

  ngOnInit(): void {
    let that = this;
    this.http.get('http://localhost:8080/api/getmovie/' + this.movieName).subscribe(function (res) {
      console.log("result of the movie", res)
      that.movie = res
    })

    this.http.get('http://localhost:8080/api/getshow/' + this.movieName).subscribe(function (res) {
      that.show = JSON.parse(JSON.stringify(res));
      that.movieShows = JSON.parse(JSON.stringify(res))
      console.log("result of the selected shows", that.movieShows)
    })

    this.http.get('http://localhost:8080/api/getcoupons/').subscribe(function (res) {
      console.log("coupons", res)
      that.coupons = res
    })

  }

  selectTicket(ticket: any, i: any, j: any, selectedShow: any) {
    let time = new Date(selectedShow.show_date).getTime() - new Date().getTime()
    this.timeDifference = time / (1000 * 3600);
    console.log("the time difference", this.timeDifference)

    let price
    if (this.timeDifference > 0 && this.timeDifference < 48) {
      price = ticket.price - (ticket.price * 5 / 100)
    } else {
      price = ticket.price
    }

    if (ticket.is_booked) {
      if (this.seatsCount < selectedShow.seatlimit) {
        console.log("seat", selectedShow.seats)
        this.selectedSeats = selectedShow.seats
        this.seatsCount = this.seatsCount + 1
        this.totalAmount = this.totalAmount + price
      } else {
        ticket.is_booked = !ticket.is_booked
        this.utilityService.snackbarmsg = "You can book maximum of " + selectedShow.seatlimit + " seats"
        this.snackbar.openFromComponent(SnackbarComponent, {
          duration: 3000
        })
      }
    } else {
      this.seatsCount = this.seatsCount - 1
      this.totalAmount = this.totalAmount - price
    }

    console.log("seats count", this.seatsCount)
  }

  getTimeDifference(show_date: any) {
    console.log("show_date", show_date)
    let time = new Date(show_date).getTime() - new Date().getTime()
    this.timeDifference = time / (1000 * 3600);
    console.log("the time difference", this.timeDifference)
    if (this.timeDifference > 0 && this.timeDifference < 48) {
      return "5% offer is there";
    } else {
      return null
    }
  }

  showTickets(selected_time: any, index: any) {
    if (this.utilityService.loggedInObj.email) {
      // console.log("this.movie.end_date", this.movie.end_date)
      if (new Date(this.selectedShow.show_date) > new Date()) {
        this.isShowTickets = true;
      } else {
        this.utilityService.snackbarmsg = "This is a past show"
        this.snackbar.openFromComponent(SnackbarComponent, {
          duration: 5000,
        })
      }
    } else {
      this.utilityService.snackbarmsg = "Please Login Before Selecting The Tickets"
      this.snackbar.openFromComponent(SnackbarComponent, {
        duration: 5000,
      })
    }

    this.selectedIndex = index
    this.selectedShow = this.show[index]
    console.log("this.selectedShow", this.selectedShow)
  }

  cancelTicketBooking() {
    this.isShowTickets = false;
    this.movieShows = JSON.parse(JSON.stringify(this.show))
  }

  bookTicket() {
    this.spinner = true
    this.userObj.email = this.utilityService.loggedInObj.email
    this.userObj.movie_name = this.selectedShow.movie_name
    this.userObj.date = this.selectedShow.show_date
    this.userObj.movie_time = this.selectedShow.movie_start_time
    let that = this;

    for (let i = 0; i < this.selectedShow.seats.length; i++) {
      for (let j = 0; j < this.selectedShow.seats[i].length; j++) {
        if (this.selectedShow.seats[i][j]["is_booked"] && this.selectedShow.seats[i][j]["email"] == this.utilityService.loggedInObj.email) {
          this.selectedShow.seats[i][j]["email"] = this.utilityService.loggedInObj.email;
          let seatNumber = this.selectedShow.seats[0].length * i + j + 1
          this.userObj.seats.push(seatNumber)
        }
      }
    }

    for (let i = 0; i < this.selectedShow.seats.length; i++) {
      for (let j = 0; j < this.selectedShow.seats[i].length; j++) {
        if (this.selectedShow.seats[i][j]["is_booked"] && this.selectedShow.seats[i][j]["email"] == '') {
          this.selectedShow.seats[i][j]["email"] = this.utilityService.loggedInObj.email;
          let seatNumber = this.selectedShow.seats[0].length * i + j + 1
          this.userObj.seats.push(seatNumber)
        }
      }
    }

    let obj = {
      email: this.userObj.email,
      my_bookings: this.userObj
    }

    console.log("book ticket", this.selectedShow)
    console.log("the user obj is: ", obj)

    this.http.put('http://localhost:8080/api/updateshow/', this.selectedShow)
      .subscribe({
        next: data => {
          if (data) {
            this.http.put('http://localhost:8080/api/updateuserbookings/', obj)
              .subscribe({
                next: user => {
                  if (user) {
                    that.utilityService.seats = that.userObj.seats
                    that.utilityService.snackbarmsg = "Ticked Booked Successfully"
                    that.snackbar.openFromComponent(SnackbarComponent, {
                      duration: 3000
                    })
                  }
                  that.spinner = false
                  console.log("updated data", data)
                  that.show[that.selectedIndex] = data
                },
                error: error => {
                  console.error('There was an error in updating the user!', error);
                  that.spinner = false
                }
              })
          }
        },
        error: error => {
          console.error('There was an error in updating the seats!', error);
          that.spinner = false
        }
      });

      this.navigateRouter.navigateByUrl("/bookings")
  }

  addCode() {
    if(this.couponPrice > 0) {
      this.utilityService.snackbarmsg = "Coupon already added"
      this.snackbar.openFromComponent(SnackbarComponent, {
        duration: 3000
      })
    } else {
      for (let i = 0; i < this.coupons.length; i++) {
        console.log("coupon matched", this.coupons[i]["CouponPrice"])
        if (this.coupons[i]["CouponCode"] == this.couponcode) {
          this.couponPrice = this.coupons[i]["CouponPrice"]
          this.totalAmount = this.totalAmount - this.couponPrice
          this.utilityService.snackbarmsg = "Coupon added successfully"
          this.snackbar.openFromComponent(SnackbarComponent, {
            duration: 3000
          })
        } else {
          this.utilityService.snackbarmsg = "Coupon does not exist"
          this.snackbar.openFromComponent(SnackbarComponent, {
            duration: 3000
          })
          this.couponPrice = 0
        }
      }
    }
    
  }

  removeCode() {
    this.couponcode = '';
    if(this.couponPrice > 0) {
      this.totalAmount = this.totalAmount + this.couponPrice 
    }
    this.couponPrice = 0;
  }

}
