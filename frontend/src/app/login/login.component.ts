import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { MatDialog, MatDialogRef } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { RegisterComponent } from '../register/register.component';
import { SnackbarComponent } from '../snackbar/snackbar.component';
import { UtilityService } from '../utility.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  logObj = {
    email: '',
    password: ''
  }

  constructor(public dialog: MatDialog, public http: HttpClient, private dialogRef: MatDialogRef<LoginComponent>, 
    public utilityService: UtilityService, public snackbar: MatSnackBar, public navigateRouter: Router) {}

  ngOnInit(): void {
  }

  login() {
    let that = this
    this.http.get('http://localhost:4200/api/login/?email=' + this.logObj.email + '&password=' + this.logObj.password)
    .subscribe(function(res){
      that.utilityService.snackbarmsg = "login successfull"
      that.snackbar.openFromComponent(SnackbarComponent, {
        duration: 3000,
      });
      that.dialogRef.close();
      console.log("result of the response", res)
      that.utilityService.loggedInObj = res;
      that.navigateRouter.navigateByUrl("/bookings")
    }, (err)=>{
      that.utilityService.snackbarmsg = "login failed, please check your credentials"
      that.snackbar.openFromComponent(SnackbarComponent, {
        duration: 3000,
      });
  })
}

  signUp() {
    const dialogRef = this.dialog.open(RegisterComponent,{
      width: '30%',
      data: null
  });

    dialogRef.afterClosed().subscribe(result => {
      console.log(`Dialog result: ${result}`);
    });
  }

}
