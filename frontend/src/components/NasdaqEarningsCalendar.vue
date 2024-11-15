<template>
  <main>

    <!-- Input type -->
    <v-container>
      <v-radio-group inline label="Please pick your input" v-model="data.inputType">
        <v-radio label="Date Range" value="range"></v-radio>
        <v-radio label="Multiple Dates" value="dates"></v-radio>
      </v-radio-group>
    </v-container>

    <!-- If date range -->
    <v-container v-if="data.inputType === 'range'">
      <v-row>
        <v-text-field type="date" label="Select Start Date" clearable v-model="data.startDate"></v-text-field>
        <v-spacer style="flex: 0 0 10px;" />
        <v-text-field type="date" label="Select End Date" clearable v-model="data.endDate"></v-text-field>
      </v-row>
    </v-container>

    <!-- If multiple dates -->
    <v-container v-if="data.inputType === 'dates'">
      <v-row>
        <v-col cols="3">
          <v-date-picker v-model="data.selectedDate" multiple range no-title
            @update:model-value="updateSelectedDates"></v-date-picker>
        </v-col>
        <v-col cols="9">
          <h3 class="date-list-title">Selected Dates</h3>
          <v-row>
            <v-col cols="12">
              <v-chip v-for="(date, index) in data.selectedDates" :key="index" class="date-chip"
                @click="removeDate(index)">
                {{ formatDate(date) }}
                <v-icon small>mdi-close</v-icon>
              </v-chip>
            </v-col>
          </v-row>
        </v-col>
      </v-row>
    </v-container>

    <!-- Submit button -->
    <v-container>
      <v-btn variant="tonal" @click="submit" :loading="data.isSubmitButtonLoading">
        Submit
      </v-btn>
    </v-container>

    <!-- Error Message -->
    <v-container v-if="data.errorMessage">
      <v-alert>{{ data.errorMessage }}</v-alert>
    </v-container>
  </main>
</template>

<script lang="ts" setup>
import { reactive } from 'vue'
import { GetNASDAQEarningsCalendar, SaveFile, WriteFile, ShowError } from '../../wailsjs/go/main/App'
import { LogInfo } from '../../wailsjs/runtime'

const data = reactive({
  inputType: 'range',
  startDate: '',
  endDate: '',
  selectedDates: [] as Date[],
  selectedDate: [] as Date[],
  isSubmitButtonLoading: false as boolean,
  errorMessage: ""
})

async function err(body: string, title: string = "") {
  return ShowError(title, body);
}

async function submit() {
  data.isSubmitButtonLoading = true
  let dates = [] as string[]
  if (data.inputType === 'range') {
    dates = getDatesInRange(data.startDate, data.endDate)
  } else if (data.inputType === 'dates') {
    dates = convertToYYYYMMDD(data.selectedDates)
  }

  let b64str = await GetNASDAQEarningsCalendar(dates)
  if (!b64str) {
    data.isSubmitButtonLoading = false
    return await err("Error Generating File")
  }
  const path = await SaveFile("NASDAQ_Earnings.xlsx");
  if(!path) {
    data.isSubmitButtonLoading = false
    return
  };
  await WriteFile(b64str, path);

  data.isSubmitButtonLoading = false
  
}

function handleExcelResponse(numArray: number[], fileName: string) {

  // Create Uint8Array from number array
  let response = new Uint8Array(numArray)

  // Create a blob object from the array of bytes
  const blob = new Blob([response], { type: 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet' });

  // Create a temporary link element
  const link = document.createElement('a');
  link.href = URL.createObjectURL(blob);
  link.setAttribute('download', fileName);
  link.style.display = 'none';

  // Add the link to the body and simulate a click
  document.body.appendChild(link);
  link.click();

  // Clean up
  document.body.removeChild(link);
  URL.revokeObjectURL(link.href);
}

function updateSelectedDates(selectedDates: Date[]) {
  data.selectedDates = selectedDates;
}
function removeDate(index: number): void {
  data.selectedDates.splice(index, 1);
}
function formatDate(date: Date): string {
  return date.toLocaleDateString(); // Format date as needed
}

function getDatesInRange(startDate: string, endDate: string): string[] {
  const datesArray: string[] = [];
  let currentDate: Date = new Date(startDate);
  const end: Date = new Date(endDate);

  while (currentDate <= end) {
    datesArray.push(currentDate.toISOString().split('T')[0]);
    currentDate.setDate(currentDate.getDate() + 1);
  }

  return datesArray;
}

function convertToYYYYMMDD(dateArray: Date[]): string[] {
  return dateArray.map((date: Date) => {
    const year = date.getFullYear();
    const month = ('0' + (date.getMonth() + 1)).slice(-2);
    const day = ('0' + date.getDate()).slice(-2);
    return `${year}-${month}-${day}`;
  });
}

</script>

<style>
main {
  padding: 1em 1em;
}

.v-radio-group>.v-input__control>.v-label {
  margin-inline-start: 0 !important;
}

.v-radio input[type="radio"] {
  opacity: 1 !important;
}

.v-radio {
  padding: 1em 1.5em 1em 0em;
}

.v-radio .v-label {
  margin-left: 0.25em;
}

.date-chip {
  margin-right: 8px;
  margin-bottom: 8px;
  /* Adjust as needed */
  /* Adjust spacing between chips if needed */
}

.date-chip v-icon {
  cursor: pointer;
  /* Add cursor pointer to the close icon */
}

.date-list-title {
  margin-bottom: 8px;
  font-weight: bold;
  font-size: 1.2em;
  /* Adjust title font size */
}
</style>
