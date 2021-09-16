import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { MatDialogRef } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { SnackbarComponent } from '../snackbar/snackbar.component';
import { UtilityService } from '../utility.service';

@Component({
  selector: 'app-addtheatres',
  templateUrl: './addtheatres.component.html',
  styleUrls: ['./addtheatres.component.css']
})
export class AddtheatresComponent implements OnInit {
  theatreName:any
  theatreImage:any

  constructor(public http: HttpClient, public utilityService: UtilityService,public snackbar: MatSnackBar, private dialogRef: MatDialogRef<AddtheatresComponent>) { }

  ngOnInit(): void {
  }

  handleFileInput(event: any) {
    this.theatreImage = event.target.files[0]
  }

  addTheatre() {
    let that = this;
    const fd = new FormData()
    fd.append("image", this.theatreImage)
    fd.append("Content-Type", "multipart/form-data")
    fd.append("theatreName", this.theatreName)
    this.http.post('http://localhost:8080/api/inserttheatre/', fd)
    .subscribe(function(res){
      console.log("result of the response", res)
      that.utilityService.snackbarmsg = res
      console.log("result of the response", res);
      that.snackbar.openFromComponent(SnackbarComponent, {
        duration: 3000,
      });
      that.dialogRef.close();
    })}
}
