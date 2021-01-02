<template>
  <v-main>
    <v-card flat>
      <v-toolbar flat>
        <v-app-bar-nav-icon @click="toggle = !toggle"></v-app-bar-nav-icon>
        <AppBar v-model="toggle"></AppBar>

        <v-img class="shrink" src="../assets/logos.png" contain height="50px" />
        <v-toolbar-title>
          {{ eventData.name }} - {{ bindinfo.module }}
        </v-toolbar-title>
        <v-toolbar-items>
          <v-card-title class="ml-2 pa-1">
            <v-chip small outlined color="deep-purple accent-4">{{
              new Date(eventData.date).toISOString().substring(0, 10)
            }}</v-chip>
          </v-card-title>
        </v-toolbar-items>

        <v-spacer></v-spacer>
        <v-btn icon v-on:click="dialog = true">
          <v-icon>mdi-plus</v-icon>
        </v-btn>
      </v-toolbar>
      <v-container fluid>
        <v-row>
          <v-col
            v-for="i in eventData.images"
            :key="i.id"
            cols="6"
            md="2"
            sm="4"
            xs="6"
          >
            <v-hover v-slot:default="{ hover }">
              <v-img class="vimage" contain aspect-ratio="1" :src="i.thumbnail">
                <v-fade-transition>
                  <v-overlay v-show="hover" absolute>
                    <v-btn icon @click.stop="deleteEntity(i)">
                      <v-icon large elevation="6" color="error">
                        mdi-delete
                      </v-icon>
                    </v-btn>
                  </v-overlay>
                </v-fade-transition></v-img
              >
            </v-hover>
          </v-col>
        </v-row>
      </v-container>
      <v-dialog v-model="dialog" persistent max-width="600px">
        <v-card>
          <v-card-title>
            <span class="headline">{{ bindinfo.module }} </span>
          </v-card-title>
          <v-form ref="form" v-model="formvalid" lazy-validation>
            <v-card-text>
              <v-text-field
                label="Description"
                outlined
                required
                v-model="description"
              ></v-text-field>

              <v-file-input
                counter
                show-size
                small-chips
                outlined
                truncate-length="5"
                label="Upload File"
                required
                multiple
                accept="image/*"
                v-model="files"
              ></v-file-input>
            </v-card-text>
          </v-form>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="blue darken-1" text @click="saveEntity()">
              Save
            </v-btn>
            <v-btn color="blue darken-1" text @click="dialog = false">
              Close
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </v-card>
    <confirmDialog ref="confirm" />
  </v-main>
</template>

<script>
import AppBar from "@/components/AppBar.vue";
export default {
  name: "AdminEntities",
  data: () => ({
    eventData: { name: "" },
    dialog: false,
    toggle: false,
    apiURL: "/api/v1/events",
    description: null,
    files: null,
    formvalid: true
  }),
  components: {
    AppBar,
    confirmDialog: () => import("@/components/ConfirmDialog.vue")
  },
  props: {
    bindinfo: Object
  },
  methods: {
    fetchData() {
      this.axios
        .get(this.apiURL + "/" + this.bindinfo.id)
        .then(response => {
          this.eventData = response.data["data"];
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
            var formData = new FormData();
            formData.append("description", this.description);
            this.files.forEach(file => {
              formData.append("files", file);
            });
            this.axios
              .post(
                this.apiURL + "/" + this.bindinfo.id + "/images",
                formData,
                {
                  headers: {
                    "Content-Type": "multipart/form-data"
                  }
                }
              )
              .then(response => {
                if (response.data.response_status.response_code == 200) {
                  this.$refs.form.reset();
                  this.dialog = false;
                  this.fetchData();
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
            this.axios
              .delete(
                this.apiURL + "/" + this.bindinfo.id + "/images/" + value.id
              )
              .then(response => {
                if (response.data.response_status.response_code == 200) {
                  this.fetchData();
                }
              });
          }
        });
    }
  },
  created() {
    this.fetchData();
  }
};
</script>

<style>
.vimage .v-image__image {
  background-size: 100% 100%;
}
</style>
