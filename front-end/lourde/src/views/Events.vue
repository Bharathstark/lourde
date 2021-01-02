<template>
  <v-main>
    <v-card flat>
      <v-toolbar flat>
        <v-app-bar-nav-icon @click="toggle = !toggle"></v-app-bar-nav-icon>
        <AppBar v-model="toggle"></AppBar>
        <v-img class="shrink" src="../assets/logos.png" contain height="50px" />
        <v-toolbar-title>
          {{ moduleName }}
        </v-toolbar-title>
        <v-spacer></v-spacer>
        <v-spacer></v-spacer>

        <v-btn icon light>
          <v-icon color="grey darken-2">
            mdi-magnify
          </v-icon>
        </v-btn>
      </v-toolbar>
      <v-container>
        <v-row fluid>
          <v-col
            v-for="i in eventData"
            :key="i.id"
            cols="6"
            md="3"
            sm="4"
            xs="6"
          >
            <v-card
              :to="module + '?id=' + i.id"
              :height="
                $vuetify.breakpoint.xs || $vuetify.breakpoint.sm ? 180 : 300
              "
              :width="
                $vuetify.breakpoint.xs || $vuetify.breakpoint.sm ? 150 : 300
              "
            >
              <v-img class="event" height="75%" :src="i.thumbnail"> </v-img>
              <v-card-title class="text-lg-h6 text-subtitle-2">
                {{ i.name }}
              </v-card-title>
            </v-card>
          </v-col>
        </v-row>
      </v-container>
    </v-card>
  </v-main>
</template>

<script>
import AppBar from "@/components/AppBar.vue";
export default {
  name: "Events",
  data: () => ({
    eventData: null,
    apiURL: null,
    toggle: false,
    moduleName: ""
  }),
  components: {
    AppBar
  },
  props: {
    module: String
  },
  methods: {
    setModule() {
      this.moduleName = this.$route.name;
    },
    loging() {},
    async fetchData() {
      this.axios
        .get("/api/v1/events?module=" + this.moduleName.toLowerCase())
        .then(response => {
          this.eventData = response.data["data"];
        })
        .catch(error => {
          console.log(error);
        });
    }
  },
  created() {
    this.loging();
    this.setModule();
    this.fetchData();
  }
};
</script>

<style>
.event > .v-image__image {
  background-size: 100% 100%;
}
</style>
