import { Component, OnInit } from "@angular/core";
import { BrewService } from "src/app/shared/services/brew/brew.service";
import {
  FormBuilder,
  FormGroup,
  FormArray,
  FormControl,
  FormControlName,
  Validators,
} from "@angular/forms";

@Component({
  selector: "brew-form",
  templateUrl: "./brew-form.component.html",
  styleUrls: ["./brew-form.component.scss"],
})
export class BrewFormComponent implements OnInit {
  inputTags: String[];
  fermentationTags: String[];
  ingredientTags: String[];
  ingredientsForm: FormGroup;
  brewForm: FormGroup;
  fermentationForm: FormGroup;
  ingredientsForm2: FormGroup;

  constructor(
    private _brewService: BrewService,
    private _formBuilder: FormBuilder
  ) {}

  ngOnInit(): void {
    this.brewForm = this._formBuilder.group({
      brewType: [""],
      Brewname: [""],
      brewDate: [""],
      mashingTemp: [""],
      mashingDuration: [""],
      totalAmount: [""],
      startDate: [""],
      endDate: [""],
    });

    this.ingredientsForm = this._formBuilder.group({
      barley: this._formBuilder.array([this._formBuilder.control("")]),
      hops: this._formBuilder.array([this._formBuilder.control("")]),
      yeast: this._formBuilder.array([this._formBuilder.control("")]),
      extra: this._formBuilder.array([this._formBuilder.control("")]),
      water: [""],
    });

    this.fermentationForm = this._formBuilder.group({
      temperature: [""],
      tempThreshold: [""],
      OG: [""],
      SG: [""],
      FG: [""],
      daysBottle: [""],
    });

    this.inputTags = [
      "Brew name",
      "Mashing temperature",
      "Brew date",
      "Mashing duration",
      "Brew type",
      "Total amount",
    ];

    this.fermentationTags = [
      "Fermentation start-date",
      "Original Gravity",
      "Fermentation end-date",
      "Specific Gravity",
      "Fermentation temperature",
      "Final Gravity",
      "Fermentation temperature threshold",
      "Bottled days",
    ];
    this.ingredientTags = ["Barley", "Hops", "Yeast", "Extra", "Water"];
  }

  get barleyForms() {
    return this.ingredientsForm.get("barley") as FormArray;
  }

  get hopsForms() {
    return this.ingredientsForm.get("hops") as FormArray;
  }

  get yeastForms() {
    return this.ingredientsForm.get("yeast") as FormArray;
  }

  get extraForms() {
    return this.ingredientsForm.get("extra") as FormArray;
  }

  addIngredient(ingredientType: string) {
    const ingredient = this._formBuilder.group({
      type: "",
      amount: "",
    });

    switch (ingredientType) {
      case "Barley":
        console.log(ingredientType);
        this.barleyForms.push(ingredient);
      case "Hops":
        this.hopsForms.push(ingredient);
      case "Yeast":
        this.yeastForms.push(ingredient);
      case "Extra":
        this.extraForms.push(ingredient);
    }
  }

  removeIngredient(ingredientType: string, i: number) {
    console.log("remove ingredient");
    switch (ingredientType) {
      case "Barley":
        console.log(ingredientType);
        this.barleyForms.removeAt(i);
      case "Hops":
        this.hopsForms.removeAt(i);
      case "Yeast":
        this.yeastForms.removeAt(i);
      case "Extra":
        this.extraForms.removeAt(i);
    }
  }

  closeForm() {
    this._brewService.toggleForm();
  }
}
