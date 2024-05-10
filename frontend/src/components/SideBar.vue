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
        </v-list>
    </v-navigation-drawer>
</template>

<script setup lang="ts">
// Add any necessary script setup code here
</script>

<style scoped>
.no-hover:hover,
.no-hover:focus {
    background-color: transparent !important;
}

.side-menu .v-list-item {
    text-align: left;
}

a {
    text-decoration: none;
    color: white;
    opacity: 0.7;
}

.square-avatar img {
    width: 36px;
    /* border-radius: 0; */
}
</style>