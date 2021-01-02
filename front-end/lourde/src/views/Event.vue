<template>
  <v-main>
    <v-card flat>
      <v-toolbar flat>
        <v-app-bar-nav-icon @click="toggle = !toggle"></v-app-bar-nav-icon>
        <AppBar v-model="toggle"></AppBar>

        <v-img class="shrink" src="../assets/logos.png" contain height="50px" />
        <v-toolbar-title>
          {{ eventData.name }}
        </v-toolbar-title>
      </v-toolbar>

      <v-tabs color="deep-purple accent-4" right>
        <v-tab>Grid</v-tab>
        <v-tab>show</v-tab>
        <v-tab-item>
          <v-container fluid>
            <v-row>
              <v-col
                v-for="i in eventData.images"
                :key="i.id"
                cols="6"
                md="3"
                sm="4"
                xs="6"
              >
                <v-img
                  class="vimage"
                  contain
                  aspect-ratio="1"
                  :src="i.thumbnail"
                ></v-img>
              </v-col>
            </v-row>
          </v-container>
        </v-tab-item>
        <v-tab-item>
          <v-carousel id="imageSlider" cycle hide-delimiter-background>
            <v-carousel-item
              contain
              v-for="i in eventData.images"
              :key="i.id"
              :src="i.thumbnail"
              reverse-transition="fade-transition"
              transition="fade-transition"
            ></v-carousel-item>
          </v-carousel>
        </v-tab-item>
      </v-tabs>
    </v-card>
  </v-main>
</template>

<script>
import AppBar from "@/components/AppBar.vue";
export default {
  name: "EventImages",
  data: () => ({
    eventData: { name: "" },
    toggle: false
  }),
  components: {
    AppBar
  },
  props: {
    eventID: String
  },
  methods: {
    async fetchData() {
      console.log(this.$route);
      this.axios
        .get("/api/v1/events/" + this.eventID)
        .then(response => {
          this.eventData = response.data["data"];
          if (response.data.response_status.response_code != 200) {
            this.$router.push("/");
          }
        })
        .catch(error => {
          console.log(error);
          this.$router.push("/");
        });
    }
  },
  created() {
    this.fetchData();
  }
};
</script>

<style>
#imageSlider {
  height: 100vh;
}
.vimage .v-image__image {
  background-size: 100% 100%;
}
</style>
