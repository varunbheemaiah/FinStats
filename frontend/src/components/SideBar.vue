<template>
    <v-navigation-drawer app>
        <router-link to="/">
            <v-list class="app-title">
                <v-list-item class="no-hover">
                    <template #prepend>
                        <v-list-item-avatar class="square-avatar">
                            <img src="../assets/images/stock-market-no-base.png" alt="FinStats Image">
                        </v-list-item-avatar>
                    </template>
                    <v-list-item-content>
                        <v-list-item-title>FinStats</v-list-item-title>
                    </v-list-item-content>
                </v-list-item>
            </v-list>
        </router-link>
        <v-divider></v-divider>
        <v-list class="side-menu">
            <!-- <router-link to="/greeting">
                <v-list-item link title="Greeting"></v-list-item>
            </router-link> -->
            <router-link to="/nasdaq-earnings">
                <v-list-item link title="NASDAQ Earnings Calendar"></v-list-item>
            </router-link>
            <router-link to="/nse-earnings">
                <v-list-item link title="NSE/BSE Earnings Calendar"></v-list-item>
            </router-link>
            <router-link to="/india-graph">
                <v-list-item link title="India Graph"></v-list-item>
            </router-link>
        </v-list>

        <template v-slot:append>
            <div class="pa-2">
                <v-btn variant="text" size="x-large" density="compact" @click="showSettings = true">
                    <v-icon>mdi-cog</v-icon>
                </v-btn>
            </div>
        </template>

        <v-dialog v-model="showSettings" max-width="500px">
            <v-card>
                <v-card-title class="headline mb-5 dialog-header-divider">Settings</v-card-title>
                <v-card-text>
                    <v-row>
                        <v-col cols="12" md="6" class="pt-16 mt-1">
                            <p>Theme</p>
                        </v-col>
                        <v-col cols="12" md="6" class="flex-center">
                            <v-radio-group v-model="chosenTheme" inline ripple @update:model-value="themeChanged">
                                <v-radio label="Light" value="light">
                                    <template #label>
                                        &nbsp;Light
                                    </template>
                                </v-radio>
                                <v-radio label="Dark" value="dark">
                                    <template #label>
                                        &nbsp;Dark
                                    </template>
                                </v-radio>
                            </v-radio-group>
                        </v-col>
                    </v-row>
                    <v-row>
                        <v-col cols="12" md="6" class="pt-16 mt-1">
                            <p>Number Format</p>
                        </v-col>
                        <v-col cols="12" md="6" class="flex-center">
                            <v-radio-group v-model="numberFormat" inline ripple @update:model-value="numberFormatChanged">
                                <v-radio label="Indian" value="indian">
                                    <template #label>
                                        &nbsp;Indian
                                    </template>
                                </v-radio>
                                <v-radio label="American" value="american">
                                    <template #label>
                                        &nbsp;American
                                    </template>
                                </v-radio>
                            </v-radio-group>
                        </v-col>
                    </v-row>
                    <!-- Add any number of settings -->
                </v-card-text>
                <v-card-actions>
                    <!-- If you want to add actions in the future, you can place them here -->
                </v-card-actions>
            </v-card>
        </v-dialog>

    </v-navigation-drawer>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useTheme } from 'vuetify';

const showSettings = ref(false);
const themer = useTheme();
const chosenTheme = ref("dark");
const numberFormat = ref("indian")

function numberFormatChanged(event: string|null): void {
    console.log(event);
}

function themeChanged(event: string|null): void {
    themer.global.name.value = chosenTheme.value;
}

</script>

<style scoped>
.no-hover:hover,
.no-hover:focus {
    background-color: transparent !important;
}

.side-menu .v-list-item {
    text-align: left;
}

.flex-center {
    display: flex;
    flex-direction: column;
    align-items: center;
}

.square-avatar img {
    width: 36px;
    /* border-radius: 0; */
}

.settings-item {
    margin-top: auto;
    cursor: pointer;
}

.settings-item .v-icon {
    margin-right: 10px;
}

.dialog-header-divider {
    border-bottom: 1px solid white;
}
</style>