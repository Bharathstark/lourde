<template>
  <v-main>
    <v-card flat>
      <v-toolbar color="white" flat>
        <v-app-bar-nav-icon @click="toggle = !toggle"></v-app-bar-nav-icon>
        <AppBar v-model="toggle"></AppBar>
        <v-img class="shrink" src="../assets/logos.png" contain height="50px" />
        <v-toolbar-title class="grey--text text--darken-4">
          Admin-Videos
        </v-toolbar-title>
        <v-spacer></v-spacer>

        <v-btn
          large
          icon
          v-on:click="(button = 'Save'), clearForm(), (videodialog = true)"
        >
          <v-icon large>mdi-plus</v-icon>
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
              <v-card-title class="ml-3 pa-0 text-lg-h6 text-subtitle-2">
                {{ i.name }}
              </v-card-title>
              <v-card-title class="ml-2 pa-0">
                <v-chip small outlined color="deep-purple accent-4">{{
                  new Date(i.date).toISOString().substring(0, 10)
                }}</v-chip>
                <v-btn icon @click.stop="editForm(i)">
                  <v-icon elevation="6" color="primary">
                    mdi-pencil
                  </v-icon>
                </v-btn>
                <v-btn icon @click.stop="deleteEntity(i)">
                  <v-icon elevation="6" color="error">
                    mdi-delete
                  </v-icon>
                </v-btn>
              </v-card-title>
            </v-card>
          </v-col>
        </v-row>
      </v-container>
    </v-card>
    <v-dialog v-model="videodialog" persistent max-width="600px">
      <v-card>
        <v-card-title>
          <span class="headline">{{ bindinfo.module }} </span>
        </v-card-title>
        <v-form ref="notiform" v-model="valid" lazy-validation>
          <v-card-text>
            <v-text-field
              label="Video Name"
              v-model="name"
              outlined
              required
            ></v-text-field>
            <v-text-field
              label="Video Description"
              outlined
              required
              v-model="description"
            ></v-text-field>
            <v-text-field
              label="Video Link"
              outlined
              required
              v-model="link"
            ></v-text-field>
            <v-menu
              v-model="datemenu"
              :close-on-content-click="false"
              :nudge-right="40"
              transition="scale-transition"
              offset-y
              min-width="290px"
            >
              <template v-slot:activator="{ on, attrs }">
                <v-text-field
                  v-model="date"
                  outlined
                  label="Video Created Date"
                  prepend-icon="mdi-calendar"
                  readonly
                  v-bind="attrs"
                  v-on="on"
                ></v-text-field>
              </template>
              <v-date-picker
                v-model="date"
                @input="datemenu = false"
              ></v-date-picker>
            </v-menu>
          </v-card-text>
        </v-form>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn
            color="blue darken-1"
            text
            @click="button == 'Save' ? saveEntity() : updateEntity()"
          >
            {{ button }}
          </v-btn>
          <v-btn color="blue darken-1" text @click="videodialog = false">
            Close
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
    <confirmDialog ref="confirm" />
  </v-main>
</template>

<script>
import AppBar from "@/components/AppBar.vue";
import confirmDialog from "@/components/ConfirmDialog.vue";
export default {
  name: "AdminVideos",
  data: () => ({
    toggle: false,
    button: "",
    name: "",
    link: "",
    date: "",
    apiURL: "/api/v1/notifications?module=videos",
    videoURL: "/api/v1/notifications",
    description: "",
    datemenu: "",
    videodialog: false,
    videosData: [],
    valid: true
  }),
  components: {
    AppBar,
    confirmDialog
  },
  props: {
    bindinfo: Object
  },
  created() {
    this.fetchData();
  },
  methods: {
    clearForm() {
      this.name = "";
      this.description = "";
      this.link = "";
      this.date = "";
    },
    fetchData() {
      this.axios
        .get(this.apiURL)
        .then(response => {
          this.videosData = response.data["data"];
        })
        .catch(error => {
          console.log(error);
        });
    },
    saveEntity() {
      this.$refs.confirm
        .open("Save", "Are you sure?", { color: "green" })
        .then(confirm => {
          if (confirm) {
            this.axios
              .post(this.apiURL, {
                name: this.name,
                url: this.link,
                description: this.description,
                date: new Date(this.date).getTime()
              })
              .then(response => {
                if (response.data.response_status.response_code == 200) {
                  this.clearForm();
                  this.videodialog = false;
                  this.fetchData();
                }
              });
          }
        });
    },
    editForm(value) {
      this.id = value.id;
      (this.name = value.name), (this.description = value.description);
      this.link = value.url;
      this.date = new Date(value.date).toISOString().substring(0, 10);
      this.button = "Update";
      this.videodialog = true;
    },
    updateEntity() {
      this.$refs.confirm
        .open("Update", "Are you sure?", { color: "blue" })
        .then(confirm => {
          if (confirm) {
            this.axios
              .put(this.videoURL + "/" + this.id, {
                name: this.name,
                url: this.link,
                date: new Date(this.date).getTime(),
                description: this.description
              })
              .then(response => {
                if (response.data.response_status.response_code == 200) {
                  this.fetchData();
                  this.clearForm();
                  this.videodialog = false;
                }
              });
          }
        });
    },
    deleteEntity(value) {
      this.$refs.confirm
        .open("Delete", "Are you sure?", { color: "red" })
        .then(confirm => {
          if (confirm) {
            this.axios.delete(this.videoURL + "/" + value.id).then(response => {
              this.id = null;
              if (response.data.response_status.response_code == 200) {
                this.fetchData();
              }
            });
          }
        });
    }
  }
};
</script>

<style></style>
