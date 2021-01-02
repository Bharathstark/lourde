<template>
  <v-main>
    <v-card flat>
      <v-toolbar flat>
        <v-app-bar-nav-icon @click="toggle = !toggle"></v-app-bar-nav-icon>
        <AppBar v-model="toggle"></AppBar>
        <v-img class="shrink" src="../assets/logos.png" contain height="50px" />
        <v-toolbar-title>
          Videos
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
        <v-row>
          <v-col
            v-for="i in videosData"
            :key="i.id"
            cols="6"
            md="3"
            sm="4"
            xs="6"
          >
            <v-card
              class="ma-6"
              :height="
                $vuetify.breakpoint.xs || $vuetify.breakpoint.sm ? 180 : 300
              "
              :width="
                $vuetify.breakpoint.xs || $vuetify.breakpoint.sm ? 150 : 300
              "
            >
              <iframe
                height="75%"
                width="100%"
                :src="'https://www.youtube-nocookie.com/embed/' + i.url"
                frameborder="0"
                gesture="media"
                allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
                allowfullscreen
              ></iframe>
              <v-card-title
                class="ml-3 pa-0 text-lg-h6 text-sm-body-1 text-subtitle-2"
              >
                {{ i.name }}
              </v-card-title>
              <v-card-title class="ml-2 pa-0">
                <v-chip
                  v-show="!($vuetify.breakpoint.xs || $vuetify.breakpoint.sm)"
                  small
                  outlined
                  color="deep-purple accent-4"
                  >{{ new Date(i.date).toISOString().substring(0, 10) }}</v-chip
                >
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
  name: "Videos",
  data: () => ({
    toggle: false,
    videosData: [],
    apiURL: "/api/v1/notifications?module=videos"
  }),
  components: {
    AppBar
  },
  methods: {
    async fetchData() {
      this.axios
        .get(this.apiURL)
        .then(response => {
          this.videosData = response.data["data"];
        })
        .catch(error => {
          console.log(error);
        });
    }
  },
  created() {
    this.fetchData();
  }
};
</script>

<style></style>
