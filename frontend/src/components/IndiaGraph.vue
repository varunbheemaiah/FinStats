<template>
    <main>
        <div class="file-input">
            <v-container>
                <v-row>
                    <v-file-input label="Upload CSV" prepend-icon="mdi-upload" accept=".csv" ref="fileInput" chips
                        @change="importCSV"></v-file-input>
                    <v-btn size="x-large" @click="showChart" class="mx-2">Generate Chart</v-btn>
                </v-row>
            </v-container>
        </div>
        <div class="chart" v-if="data.showChart">
            <apexchart type="line" :options="data.chartConfig.options" :series="data.chartConfig.series" height="810">
            </apexchart>
        </div>
    </main>
</template>

<script lang="ts" setup>
import { reactive, ref, watch } from 'vue'
import { useTheme } from 'vuetify';

const themer = useTheme()
const data = reactive({
    csv: "" as String | ArrayBuffer | null,
    showChart: false as boolean,
    chartConfig: {} as any
})
const fileInput = ref<HTMLInputElement | null>(null)
const files = ref()

function importCSV() {
    files.value = fileInput.value?.files

    if (files.value.length === 0) return

    let file = files.value[0] as File

    if (!file.name.endsWith(".csv") && file.type !== "text/csv") return

    let reader = new FileReader()
    reader.readAsText(file)
    reader.onload = () => {
        data.csv = reader.result
    }

}

function cleanCSV(csv: string): string {
    // Regular expression to find commas inside quotes
    const regex = /"([^"]*)"/g;

    // Replace commas inside quotes with an empty string
    let result = csv.replace(regex, (match) => {
        return match.replace(/,/g, '');
    });

    result = result.replace(/['"]+/g, '')

    return result;
}

function csvToJSON(csv: string): Record<string, string>[] {
    const lines = csv.split("\n");
    if (lines.length === 0) {
        return [];
    }
    const result: Record<string, string>[] = [];
    const headers = lines[0].split(",").map(header => header.trim());
    for (let i = 1; i < lines.length; i++) {
        const line = lines[i].trim();
        if (!line) {
            continue;
        }
        const words = line.split(",");
        const obj: Record<string, string> = {};
        for (let j = 0; j < headers.length; j++) {
            obj[headers[j]] = words[j] || '';
        }
        result.push(obj);
    }

    return result;
}

function convertStringToDate(dateString: string): Date {
    const months: string[] = [
        "Jan", "Feb", "Mar", "Apr", "May", "Jun",
        "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"
    ];

    // Split the date string into components
    const [day, month, year] = dateString.split('-');

    // Convert the month string to a month index
    const monthIndex: number = months.indexOf(month);

    if (monthIndex === -1) {
        throw new Error("Invalid month in date string");
    }

    // Create and return the Date object
    return new Date(parseInt(year, 10), monthIndex, parseInt(day, 10));
}

function numberWithCommas(x: number) {

    // INDIAN
    const [integerPart, fractionalPart] = x.toString().split('.');
    let str = integerPart;
    let lastThree = str.slice(-3);
    let otherNumbers = str.slice(0, -3);
    if (otherNumbers !== '') {
        lastThree = ',' + lastThree;
    }
    const formattedInteger = otherNumbers.replace(/\B(?=(\d{2})+(?!\d))/g, ",") + lastThree;
    return fractionalPart ? formattedInteger + '.' + fractionalPart : formattedInteger;

    // AMERICAN
    // return x.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",");
}

watch(themer.global.name, (newValue) => {
    if(data.showChart) {
        showChart()
    }
})

function showChart() {

    if (!data.csv) return
    let cleanedCsv = cleanCSV(data.csv as string)
    let chartData = csvToJSON(cleanedCsv)
    chartData.reverse()

    let labels = chartData.map(x => x.Date)
    let series = [
        {
            name: "Open",
            type: "line",
            data: chartData.map(x => parseFloat(x.OPEN))
        },
        {
            name: "Close",
            type: "line",
            data: chartData.map(x => parseFloat(x.close))
        },
        {
            name: "High",
            type: "line",
            data: chartData.map(x => parseFloat(x.HIGH))
        },
        {
            name: "Low",
            type: "line",
            data: chartData.map(x => parseFloat(x.LOW))
        },
        {
            name: "Volume",
            type: "column",
            data: chartData.map(x => parseInt(x.VOLUME))
        },
        {
            name: "No. of Trades",
            type: "column",
            data: chartData.map(x => parseInt(x['No of trades']))
        },
    ]

    let max = -Infinity
    for (let i = 0; i < 4; i++) {
        const dataMax = Math.max(...series[i].data)
        if (dataMax > max) max = dataMax
    }

    data.chartConfig = {
        options: {
            chart: {
                id: 'price-chart'
            },
            xaxis: {
                categories: labels,
                tickAmount: 12,
            },
            colors: ['#008FFB', '#FEB019', '#00E396', '#FF4560', '#458B74', '#775DD0'],
            yaxis: [
                {
                    show: true,
                    showAlways: true,
                    min: 0,
                    max: max,
                    title: {
                        text: 'Value'
                    },
                    labels: {
                        formatter: function (value: number) {
                            return numberWithCommas(value);
                        }
                    }
                },
                {
                    show: false,
                    showAlways: true,
                    min: 0,
                    max: max,
                    title: {
                        text: 'Value'
                    },
                    labels: {
                        formatter: function (value: number) {
                            return numberWithCommas(value);
                        }
                    }
                },
                {
                    show: false,
                    showAlways: true,
                    min: 0,
                    max: max,
                    title: {
                        text: 'Value'
                    },
                    labels: {
                        formatter: function (value: number) {
                            return numberWithCommas(value);
                        }
                    }
                },
                {
                    show: false,
                    showAlways: true,
                    min: 0,
                    max: max,
                    title: {
                        text: 'Value'
                    },
                    labels: {
                        formatter: function (value: number) {
                            return numberWithCommas(value);
                        }
                    }
                },
                {
                    show: true,
                    showAlways: true,
                    opposite: true,
                    title: {
                        text: 'Volume'
                    },
                    labels: {
                        formatter: function (value: number) {
                            return numberWithCommas(value);
                        }
                    }
                },
                {
                    show: false,
                    showAlways: true,
                    title: {
                        text: 'Value'
                    },
                    labels: {
                        formatter: function (value: number) {
                            return numberWithCommas(value);
                        }
                    }
                },
            ],
            stroke: {
                show: true,
                curve: 'straight',
                lineCap: 'butt',
                colors: undefined,
                width: 1,
                dashArray: 0,
            },
            theme: {
                mode: themer.global.name.value
            },
        },
        series: series,
    }

    data.showChart = true

}

</script>
