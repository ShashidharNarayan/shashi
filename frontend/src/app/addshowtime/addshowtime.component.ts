import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { MatDialogRef } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { SnackbarComponent } from '../snackbar/snackbar.component';
import { UtilityService } from '../utility.service';

@Component({
  selector: 'app-addshowtime',
  templateUrl: './addshowtime.component.html',
  styleUrls: ['./addshowtime.component.css']
})
export class AddshowtimeComponent implements OnInit {
  
  constructor(public utilityService: UtilityService, public http: HttpClient, public snackbar: MatSnackBar, private dialogRef: MatDialogRef<AddshowtimeComponent>) {
   
   }

  showobj = {
    theatre_name: '',
    movie_name: '',
    show_date: new Date(),
    start_time: {hour: 0, minute: 0},
    silver_category_seats: '',
    gold_category_seats: '',
    silver_category_price: '',
    gold_category_price: '',
    movie_start_time: '',
    seats: '',
    seatlimit: ''
  }

  ngOnInit(): void {
  }

  // toTime(timeString: any){ 
  //   const [time, modifier] = timeString.split(' ');
  //   let [hours, minutes] = time.split(':');
  
  //   if (hours === '12') {
  //     hours = '00';
  //   }
  
  //   if (modifier === 'PM') {
  //     hours = parseInt(hours, 10) + 12;
  //   }
  
  //   return `${hours}:${minutes}`;
  // }

  addShow() {
    let gold_category_seats = parseInt(this.showobj.gold_category_seats)
    let silver_category_seats = parseInt(this.showobj.silver_category_seats)
    let seatsArr = []
    let arr:any = []
    let chunksize = 10;

    for(let i=0; i < silver_category_seats; i++) {
      seatsArr.push({'is_booked': false, 'email': '', 'price': this.showobj.silver_category_price, 'category': 'silver'})
    }

    for(let i=0; i < gold_category_seats; i++) {
      seatsArr.push({'is_booked': false, 'email': '', 'price': this.showobj.gold_category_price, 'category': 'gold'})
    }

    seatsArr.forEach((item)=>{
      if(!arr.length || arr[arr.length-1].length == chunksize)
      arr.push([]);
    
      arr[arr.length-1].push(item);
    });

    // console.log("show obj", this.showobj)

    this.showobj.seats = arr;
    let minutes;

    if(this.showobj.start_time.minute < 10) {
      minutes = "0" + this.showobj.start_time.minute
    } else {
      minutes = this.showobj.start_time.minute
    }

    this.showobj.movie_start_time = this.showobj.start_time.hour + ": " + minutes
    let that = this;

    // adding 5:30 hrs for time conversion
    this.showobj.show_date = new Date(new Date(this.showobj.show_date).setHours(this.showobj.start_time.hour + 5, this.showobj.start_time.minute + 30, 0));
    // this.showobj.end_date = new Date(new Date(this.showobj.end_date).setHours(this.showobj.end_time.hour + 5,this.showobj.end_time.minute + 30, 0))
    console.log("the date object is", this.showobj)
    // this.showobj.start_time = this.showobj.start_time.hour + ":" + this.showobj.start_time.minute 
    this.http.post('http://localhost:8080/api/insertshow/', this.showobj)
    .subscribe(function(res){
      that.utilityService.snackbarmsg = res
      console.log("result of the response", res);
      that.snackbar.openFromComponent(SnackbarComponent, {
        duration: 3000,
      });
      that.dialogRef.close();
    })
  }

}
