<template>
  <v-main>
    <v-toolbar color="white" flat>
      <v-app-bar-nav-icon @click="toggle = !toggle"></v-app-bar-nav-icon>
      <AppBar v-model="toggle"></AppBar>
      <v-img class="shrink" src="../assets/logos.png" contain height="50px" />
      <v-toolbar-title class="grey--text text--darken-4">
        DashBoard
      </v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn text color="primary" v-on:click="changePassword">
        Change Password
      </v-btn>
      <v-btn text color="primary" v-on:click="logout">
        logout
      </v-btn>
    </v-toolbar>
    <v-container id="dashboard" fluid tag="section">
      <v-row>
        <v-col cols="12" sm="6" lg="3">
          <MaterialStatsCard
            color="info"
            icon="mdi-twitter"
            title="Server Status"
            value="OK"
            sub-icon="mdi-clock"
            sub-text="Just Updated"
          />
        </v-col>
        <v-col cols="12" sm="6" lg="3">
          <MaterialStatsCard
            color="info"
            icon="mdi-twitter"
            title="DB status"
            value="OK"
            sub-icon="mdi-clock"
            sub-text="Just Updated"
          />
        </v-col>
        <v-col cols="12" sm="6" lg="3">
          <MaterialStatsCard
            color="info"
            icon="mdi-calendar-range-outline"
            title="Events"
            :value="healthData.events"
            sub-icon="mdi-clock"
            sub-text="Events data"
          />
        </v-col>
        <v-col cols="12" sm="6" lg="3">
          <MaterialStatsCard
            color="success"
            icon="mdi-trophy"
            title="Achievements"
            :value="healthData.achievements"
            sub-icon="mdi-trophy"
            sub-text="Just Updated"
          />
        </v-col>

        <v-col cols="12" sm="6" lg="3">
          <MaterialStatsCard
            color="orange"
            icon="mdi-bell"
            title="Notifications"
            :value="healthData.notification"
            sub-icon="mdi-bell"
            sub-icon-color="red"
            sub-text="Get More Space..."
          />
        </v-col>
        <v-col cols="12" sm="6" lg="3">
          <MaterialStatsCard
            color="orange"
            icon="mdi-youtube"
            title="Videos"
            :value="healthData.videos"
            sub-icon="mdi-youtube"
            sub-icon-color="red"
            sub-text="Youtube Videos"
          />
        </v-col>
        <v-col cols="12" sm="6" lg="3">
          <MaterialStatsCard
            color="orange"
            icon="mdi-account-group-outline"
            title="Students"
            :value="healthData.academics"
            sub-icon="mdi-account-group-outline"
            sub-icon-color="blue"
            sub-text="Academic result Entries"
          />
        </v-col>
        <v-col cols="12" sm="6" lg="3">
          <MaterialStatsCard
            color="primary"
            icon="mdi-calendar"
            title="Website Visits"
            :value="healthData.calendar"
            sub-icon="mdi-calendar"
            sub-text="Tracked from Google Analytics"
          />
        </v-col>
      </v-row>

      <v-row fluid>
        <v-col cols="12" md="6" sm="6">
          <MaterialCard class="px-5 py-3 ma-10">
            <template v-slot:heading>
              <div class="display-2 font-weight-light">
                Notifications
              </div>

              <div class="subtitle-1 font-weight-light">
                Notifications to be displayed
              </div>
            </template>

            <v-list class="overflow-y-auto" height="300">
              <v-list-item
                v-for="notification in notiData.notification"
                :key="notification.id"
              >
                <v-list-item-content class="wrap-text">
                  {{ notification.name }}
                </v-list-item-content>
                <v-list-item-icon @click="editForm(notification)">
                  <v-icon>mdi-pencil</v-icon>
                </v-list-item-icon>
                <v-list-item-icon @click="deleteEntity(notification)">
                  <v-icon>mdi-delete</v-icon>
                </v-list-item-icon>
              </v-list-item>
            </v-list>
            <template v-slot:actions>
              <v-btn
                fab
                @click="(notidialog = true), (button = 'Save'), clearForm()"
                absolute
                bottom
                right
              >
                <v-icon>mdi-plus</v-icon>
              </v-btn>
            </template>
          </MaterialCard>
        </v-col>
        <v-col cols="6">
          <MaterialCard class="px-5 py-3 ma-10">
            <template v-slot:heading>
              <div class="display-2 font-weight-light">
                Banner
              </div>
              <div class="subtitle-1 font-weight-light">
                Banner to be displayed
              </div>
            </template>
            <v-card flat>
              <v-form ref="bannerForm">
                <v-text-field
                  v-model="bannername"
                  outlined
                  label="Banner"
                ></v-text-field>
                <v-text-field
                  v-model="bannerlink"
                  outlined
                  label="Link"
                ></v-text-field>

                <v-switch
                  v-model="banneractive"
                  value="active"
                  label="banner"
                ></v-switch>

                <v-btn
                  class="align-end"
                  color="blue darken-1"
                  text
                  @click="
                    editBanner(bannerid, bannername, bannerlink, banneractive)
                  "
                >
                  Save
                </v-btn>
              </v-form>
            </v-card>
          </MaterialCard>
        </v-col>
      </v-row>
    </v-container>
    <v-dialog v-model="notidialog" persistent max-width="600px">
      <v-card>
        <v-card-title>
          <span class="headline">{{ bindinfo.module }} </span>
        </v-card-title>
        <v-form ref="notiform" v-model="valid" lazy-validation>
          <v-card-text>
            <v-text-field
              label="Notification"
              v-model="name"
              outlined
              required
            ></v-text-field>
            <v-text-field
              label="Link"
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
                  label="Created Date"
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
          <v-btn color="blue darken-1" text @click="notidialog = false">
            Close
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
    <v-dialog v-model="chanPass" persistent max-width="600px">
      <v-card>
        <v-card-title>
          <span class="headline">changePassword</span>
        </v-card-title>
        <v-form ref="passWord" v-model="valid" lazy-validation>
          <v-card-text>
            <v-text-field
              label="Old Password"
              v-model="oldPassword"
              type="password"
              :rules="[v => !!v || 'Item is required']"
              outlined
              required
            ></v-text-field>
            <v-text-field
              label="New Password"
              outlined
              required
              type="password"
              :rules="[v => !!v || 'Item is required']"
              v-model="newPassword"
            ></v-text-field>
            <v-text-field
              label="Confirm Password"
              outlined
              required
              :value="newPassword"
            ></v-text-field>
          </v-card-text>
        </v-form>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn
            color="blue darken-1"
            text
            :loading="loading"
            @click="savePassword"
          >
            Update
          </v-btn>
          <v-btn
            color="blue darken-1"
            text
            @click="(chanPass = false), $refs.passWord.reset()"
          >
            Close
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
    <confirmDialog ref="confirm" />
  </v-main>
</template>

<script>
import MaterialCard from "@/components/MaterialCard.vue";
import MaterialStatsCard from "@/components/MaterialStatsCard.vue";
import AppBar from "@/components/AppBar.vue";
import confirmDialog from "@/components/ConfirmDialog.vue";
export default {
  name: "DashBoard",
  components: {
    MaterialCard,
    MaterialStatsCard,
    AppBar,
    confirmDialog
  },
  data: () => ({
    toggle: false,
    notidialog: false,
    oldPassword: "",
    newPassword: "",
    valid: false,
    datemenu: false,
    loading: false,
    notiURL: "/api/v1/notifications?module=notification",
    bannerURL: "/api/v1/notifications?module=banner ",
    apiURL: "/api/v1/notifications",
    changeUrl: "/api/v1/change",
    healthAPI: "/api/v1/health",
    logoutUrl: "/api/auth/logout",
    healthData: {},
    notiData: { notification: [], banner: {} },
    bannername: "",
    bannerlink: "",
    banneractive: "",
    chanPass: false,
    bannerid: "",
    button: "",
    name: "",
    link: "",
    date: "",
    banner: "",
    items: [
      {
        icon: "mdi-wifi",
        text: "Wifi bhsdbhsdhbshd sdbsbdbsd"
      },
      {
        icon: "mdi-bluetooth",
        text: "Wifi bhsdbhsdhbshd sdbsbdbsd"
      },
      {
        icon: "mdi-chart-donut",
        text:
          "Wifi bhsdbhsdhbshd sdbsbdbsd ndjndnjsndjnjsddddddddddddddddd ddddddddddddddddddddddddddddddddddddddddd"
      },
      {
        icon: "mdi-chart-donut",
        text:
          "Wifi bhsdbhsdhbshd sdbsbdbsd ndjndnjsndjnjsdddddddddddddddd dddddddddddddddddddddddddddddddddddddddddd"
      },
      {
        icon: "mdi-chart-donut",
        text:
          "Wifi bhsdbhsdhbshd sdbsbdbsd ndjndnjsndjnjsddddddddddddddc dddddddddddddddddddddddddddddddddddddddddddd"
      },
      {
        icon: "mdi-chart-donut",
        text:
          "Wifi bhsdbhsdhbshd sdbsbdbsd ndjndnjsndjnjsdddddddddddddddddddddddddddddddddddddddddddddddddddddddddd"
      }
    ],
    model: 1
  }),
  props: {
    bindinfo: Object
  },
  watch: {
    "notiData.banner": function(newValue) {
      this.bannername = newValue.name;
      this.bannerlink = newValue.url;
      this.banneractive = newValue.description;
      this.bannerid = newValue.id;
    }
  },
  created() {
    this.fetchNotiData("notification", this.notiURL);
    this.fetchNotiData("banner", this.bannerURL);
    this.fetchHelathData();
  },
  methods: {
    logout() {
      this.axios
        .get(this.logoutUrl)
        .then(response => {
          if (response.data.response_status.response_code == 200) {
            this.$cookies.remove("csrf_token");
            this.$router.push("/login");
          }
        })
        .catch(error => {
          console.log(error);
        });
    },
    fetchHelathData() {
      this.axios
        .get(this.healthAPI)
        .then(response => {
          this.healthData = response.data["data"];
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
              .post(this.notiURL, {
                name: this.name,
                url: this.link,
                date: new Date(this.date).getTime()
              })
              .then(response => {
                if (response.data.response_status.response_code == 200) {
                  this.$refs.notiform.reset();
                  this.notidialog = false;
                  this.fetchNotiData("notification", this.notiURL);
                }
              });
          }
        });
    },
    changePassword() {
      this.chanPass = true;
    },
    fetchNotiData(key, url) {
      this.axios
        .get(url)
        .then(response => {
          if (key == "banner") {
            this.notiData[key] = response.data["data"][0];
          } else {
            this.notiData[key] = response.data["data"];
          }
        })
        .catch(error => {
          console.log(error);
        });
    },
    updateEntity() {
      this.$refs.confirm
        .open("Update", "Are you sure?", { color: "blue" })
        .then(confirm => {
          if (confirm) {
            this.axios
              .put(this.apiURL + "/" + this.id, {
                name: this.name,
                url: this.link,
                date: new Date(this.date).getTime(),
                description: this.description
              })
              .then(response => {
                if (response.data.response_status.response_code == 200) {
                  this.clearForm();
                  this.notidialog = false;
                  this.fetchNotiData("notification", this.notiURL);
                  this.fetchNotiData("banner", this.bannerURL);
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
            this.axios.delete(this.apiURL + "/" + value.id).then(response => {
              this.id = null;
              if (response.data.response_status.response_code == 200) {
                this.fetchNotiData("notification", this.notiURL);
              }
            });
          }
        });
    },
    clearForm() {
      this.name = "";
      this.description = "";
      this.link = "";
      this.date = "";
      this.colorpick = "";
    },
    savePassword() {
      this.loading = true;
      const params = new URLSearchParams();
      params.append("username", this.$cookies.get("username"));
      params.append("oldpassword", this.oldPassword);
      params.append("newpassword", this.newPassword);
      this.axios
        .post(this.changeUrl, params, {
          headers: {
            "Content-Type": "application/x-www-form-urlencoded"
          }
        })
        .then(response => {
          if (response.data.response_status.response_code == 200) {
            this.loading = false;
            this.$parent.$refs.showed.open(
              "Password Change Successfull - kindly Relogin",
              "blue",
              -1,
              true
            );
          } else {
            this.loading = false;
            this.$parent.$refs.showed.open(
              "Password Change Failed - User Name Password Error",
              "red",
              2000,
              false
            );
          }
        });
    },
    editForm(value) {
      this.id = value.id;
      this.name = value.name;
      this.link = value.url;
      this.date = new Date(value.date).toISOString().substring(0, 10);
      this.button = "Update";
      this.notidialog = true;
    },
    editBanner(id, banner, link, active) {
      this.id = id;
      this.name = banner;
      this.link = link;
      this.description = active == "active" ? "active" : "inactive";
      this.date = 0;
      this.updateEntity();
    }
  }
};
</script>

<style></style>
