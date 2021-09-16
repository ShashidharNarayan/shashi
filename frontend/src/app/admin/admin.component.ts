import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { AddcouponsComponent } from '../addcoupons/addcoupons.component';
import { AddmovieComponent } from '../addmovie/addmovie.component';
import { AddshowtimeComponent } from '../addshowtime/addshowtime.component';
import { AddtheatresComponent } from '../addtheatres/addtheatres.component';
import { SnackbarComponent } from '../snackbar/snackbar.component';
import { UtilityService } from '../utility.service';

@Component({
  selector: 'app-admin',
  templateUrl: './admin.component.html',
  styleUrls: ['./admin.component.css']
})
export class AdminComponent implements OnInit {
  selectedCategory: any = "movie"
  email: any
  is_add_admin: any = false

  constructor(public dialog: MatDialog, public http: HttpClient, public utilityService: UtilityService, private snackbar: MatSnackBar) { }
  
  ngOnInit(): void {
  }

  openAddMovieDialog() {
    const dialogRef = this.dialog.open(AddmovieComponent, {
      width: '30%',
      data: null
    });

    dialogRef.afterClosed().subscribe(result => {
      console.log(`Dialog result: ${result}`);
    });
  }

  openAddShowTimeDialog() {
    const dialogRef = this.dialog.open(AddshowtimeComponent, {
      width: '45%',
      data: null
    });

    dialogRef.afterClosed().subscribe(result => {
      console.log(`Dialog result: ${result}`);
    });
  }

  openAddTheatreDialog() {
    const dialogRef = this.dialog.open(AddtheatresComponent, {
      width: '30%',
      data: null
    });

    dialogRef.afterClosed().subscribe(result => {
      console.log(`Dialog result: ${result}`);
    });
  }

  openAddCouponsDialog() {
    const dialogRef = this.dialog.open(AddcouponsComponent, {
      width: '30%',
      data: null
    });

    dialogRef.afterClosed().subscribe(result => {
      console.log(`Dialog result: ${result}`);
    });
  }

  deleteMovie(idx: any) {
    let id = this.utilityService.movies[idx]['id']
    console.log("the id of the movie", id)
    this.http.delete('http://localhost:8080/api/delete/'+ id)
      .subscribe(
        result => console.log(result),
        err => console.error(err)
      );
  }

  addAdmin() {
    console.log("email id is:", this.email)
    let that = this
    this.http.put('http://localhost:8080/api/updateuserasadmin/', {email: this.email, is_admin: true})
              .subscribe({
                next: user => {
                  if (user) {
                    that.utilityService.snackbarmsg = "user updated"
                    that.snackbar.openFromComponent(SnackbarComponent, {
                      duration: 3000
                    })
                  }
                },
                error: error => {
                  console.error('There was an error in updating the user!', error);
                }
              })
              this.is_add_admin = false
              this.email = ""
  }


}
