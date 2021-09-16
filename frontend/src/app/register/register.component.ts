import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { MatDialogRef } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { SnackbarComponent } from '../snackbar/snackbar.component';
import { UtilityService } from '../utility.service';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit {
  spinner: boolean = false
  regObj = {
    name: '',
    gender: '',
    email: '',
    password: '',
    confirmPassword: ''
  }

  constructor(public http: HttpClient, public snackbar: MatSnackBar, public utilityService: UtilityService, private dialogRef: MatDialogRef<RegisterComponent>) { }

  ngOnInit(): void {
  }

  signUp() {
    this.spinner = true;
    console.log("signup details", this.regObj)
    let that = this;
    if (this.regObj.password == this.regObj.confirmPassword) {
      this.http.post('http://localhost:8080/api/register/', this.regObj)
        .subscribe(function (res) {
          if (res == 'user already exists') {
            that.utilityService.snackbarmsg = "User already exists"
          } else {
            console.log("registration result", res)
            that.utilityService.loggedInObj = res
            that.utilityService.snackbarmsg = res//"You have registered successfully"
          }

          that.snackbar.openFromComponent(SnackbarComponent, {
            duration: 3000,
          });
        }, (err) => {
          console.log("err in registration", err)
          that.utilityService.snackbarmsg = err.error//"Registration failed, please try again after some time"
          that.snackbar.openFromComponent(SnackbarComponent, {
            duration: 3000,
          });
        })
    } else {
      that.utilityService.snackbarmsg = "passwords are not matching .."
      that.snackbar.openFromComponent(SnackbarComponent, {
        duration: 3000,
        data: "passwords are not matching .."
      });
      console.log("passwords are not matching ..")
    }
    this.spinner = false
    this.dialogRef.close();
  }

}
