<template>
  <v-main>
    <v-toolbar color="white" flat>
      <v-app-bar-nav-icon @click="toggle = !toggle"></v-app-bar-nav-icon>
      <AppBar v-model="toggle"></AppBar>

      <v-img class="shrink" src="../assets/logos.png" contain height="50px" />
      <v-toolbar-title class="grey--text text--darken-4">
        Admin-Calendar
      </v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn large icon v-on:click="saveForm">
        <v-icon large>mdi-plus</v-icon>
      </v-btn>
    </v-toolbar>

    <v-row>
      <v-col>
        <v-sheet height="64">
          <v-toolbar flat>
            <v-btn
              outlined
              class="mr-4"
              color="grey darken-2"
              @click="setToday"
            >
              Today
            </v-btn>
            <v-btn fab text small color="grey darken-2" @click="prev">
              <v-icon small>
                mdi-chevron-left
              </v-icon>
            </v-btn>
            <v-btn fab text small color="grey darken-2" @click="next">
              <v-icon small>
                mdi-chevron-right
              </v-icon>
            </v-btn>
            <v-toolbar-title v-if="$refs.calendar">
              {{ $refs.calendar.title }}
            </v-toolbar-title>
            <v-spacer></v-spacer>
            <v-menu bottom right>
              <template v-slot:activator="{ on, attrs }">
                <v-btn outlined color="grey darken-2" v-bind="attrs" v-on="on">
                  <span>{{ typeToLabel[type] }}</span>
                  <v-icon right>
                    mdi-menu-down
                  </v-icon>
                </v-btn>
              </template>
              <v-list>
                <v-list-item @click="type = 'day'">
                  <v-list-item-title>Day</v-list-item-title>
                </v-list-item>
                <v-list-item @click="type = 'week'">
                  <v-list-item-title>Week</v-list-item-title>
                </v-list-item>
                <v-list-item @click="type = 'month'">
                  <v-list-item-title>Month</v-list-item-title>
                </v-list-item>
                <v-list-item @click="type = '4day'">
                  <v-list-item-title>4 days</v-list-item-title>
                </v-list-item>
              </v-list>
            </v-menu>
          </v-toolbar>
        </v-sheet>
        <v-sheet height="600">
          <v-calendar
            ref="calendar"
            v-model="focus"
            color="primary"
            :events="eventData"
            :type="type"
            :event-color="getEventColor"
            @click:event="showEvent"
            @click:more="viewDay"
            @click:date="viewDay"
          ></v-calendar>
          <v-menu
            v-model="selectedOpen"
            :close-on-content-click="false"
            :activator="selectedElement"
            offset-x
          >
            <v-card color="grey lighten-4" min-width="350px" flat>
              <v-toolbar :color="selectedEvent.color" dark>
                <v-toolbar-title v-html="selectedEvent.name"></v-toolbar-title>
                <v-spacer></v-spacer>
                <v-btn
                  v-if="selectedEvent.link != ''"
                  icon
                  :href="selectedEvent.link"
                  target="_blank"
                >
                  <v-icon>mdi-open-in-new</v-icon>
                </v-btn>
                <v-btn icon @click.stop="editForm(selectedEvent)">
                  <v-icon>mdi-pencil</v-icon>
                </v-btn>
                <v-btn icon @click="deleteEntity(selectedEvent)">
                  <v-icon>mdi-delete</v-icon>
                </v-btn>
              </v-toolbar>
              <v-card-text class="text-wrap">
                <span v-html="selectedEvent.description"></span>
              </v-card-text>
              <v-card-actions>
                <v-btn text color="secondary" @click="selectedOpen = false">
                  Cancel
                </v-btn>
              </v-card-actions>
            </v-card>
          </v-menu>
        </v-sheet>
      </v-col>
    </v-row>
    <v-dialog v-model="dialog" persistent max-width="600px">
      <v-card>
        <v-card-title>
          <span class="headline">Calendar </span>
        </v-card-title>
        <v-form ref="form" v-model="valid" lazy-validation>
          <v-card-text>
            <v-text-field
              label="Name"
              v-model="name"
              :rules="[v => !!v || 'field is required']"
              outlined
              required
            ></v-text-field>

            <v-text-field
              label="Description"
              outlined
              required
              v-model="description"
            ></v-text-field>
            <v-text-field
              label="Link"
              outlined
              required
              v-model="link"
            ></v-text-field>
            <v-row>
              <v-col cols="12" md="6">
                <v-menu
                  v-model="startmenu"
                  :close-on-content-click="false"
                  :nudge-right="40"
                  transition="scale-transition"
                  offset-y
                  min-width="290px"
                >
                  <template v-slot:activator="{ on, attrs }">
                    <v-text-field
                      v-model="startdate"
                      :rules="[v => !!v || 'field is required']"
                      outlined
                      label="Start Date"
                      prepend-icon="mdi-calendar"
                      readonly
                      v-bind="attrs"
                      v-on="on"
                    ></v-text-field>
                  </template>
                  <v-date-picker
                    v-model="startdate"
                    @input="startmenu = false"
                  ></v-date-picker>
                </v-menu>
              </v-col>
              <v-col cols="12" md="6">
                <v-dialog
                  ref="starttimedialog"
                  v-model="modal2"
                  :return-value.sync="starttime"
                  persistent
                  width="290px"
                >
                  <template v-slot:activator="{ on, attrs }">
                    <v-text-field
                      v-model="starttime"
                      :rules="[v => !!v || 'field is required']"
                      label="Start Time"
                      prepend-icon="mdi-clock-time-four-outline"
                      outlined
                      readonly
                      v-bind="attrs"
                      v-on="on"
                    ></v-text-field>
                  </template>
                  <v-time-picker
                    v-if="modal2"
                    v-model="starttime"
                    full-width
                    format="ampm"
                    use-seconds
                  >
                    <v-spacer></v-spacer>
                    <v-btn text color="primary" @click="modal2 = false">
                      Cancel
                    </v-btn>
                    <v-btn
                      text
                      color="primary"
                      @click="$refs.starttimedialog.save(starttime)"
                    >
                      OK
                    </v-btn>
                  </v-time-picker>
                </v-dialog>
              </v-col>
            </v-row>
            <v-row>
              <v-col cols="12" md="6">
                <v-menu
                  v-model="endmenu"
                  :close-on-content-click="false"
                  :nudge-right="40"
                  transition="scale-transition"
                  offset-y
                  min-width="290px"
                >
                  <template v-slot:activator="{ on, attrs }">
                    <v-text-field
                      v-model="enddate"
                      outlined
                      :rules="[v => !!v || 'field is required']"
                      label="End Date"
                      prepend-icon="mdi-calendar"
                      readonly
                      v-bind="attrs"
                      v-on="on"
                    ></v-text-field>
                  </template>
                  <v-date-picker
                    v-model="enddate"
                    @input="endmenu = false"
                  ></v-date-picker>
                </v-menu>
              </v-col>
              <v-col cols="12" md="6">
                <v-dialog
                  ref="endtimedialog"
                  v-model="modal3"
                  :return-value.sync="endtime"
                  persistent
                  width="290px"
                >
                  <template v-slot:activator="{ on, attrs }">
                    <v-text-field
                      v-model="endtime"
                      label="End Time"
                      :rules="[v => !!v || 'field is required']"
                      prepend-icon="mdi-clock-time-four-outline"
                      outlined
                      readonly
                      v-bind="attrs"
                      v-on="on"
                    ></v-text-field>
                  </template>
                  <v-time-picker
                    v-if="modal3"
                    format="ampm"
                    use-seconds
                    v-model="endtime"
                    full-width
                  >
                    <v-spacer></v-spacer>
                    <v-btn text color="primary" @click="modal3 = false">
                      Cancel
                    </v-btn>
                    <v-btn
                      text
                      color="primary"
                      @click="$refs.endtimedialog.save(endtime)"
                    >
                      OK
                    </v-btn>
                  </v-time-picker>
                </v-dialog>
              </v-col>
            </v-row>
            <v-menu
              v-model="colormenu"
              :close-on-content-click="false"
              :nudge-right="40"
              transition="scale-transition"
              offset-y
              min-width="290px"
            >
              <template v-slot:activator="{ on, attrs }">
                <v-text-field
                  v-model="colorpick"
                  outlined
                  label="Pick Color"
                  prepend-icon="mdi-calendar"
                  readonly
                  @input="colormenu = true"
                  v-bind="attrs"
                  v-on="on"
                ></v-text-field>
              </template>
              <v-color-picker
                v-model="colorpick"
                dot-size="5"
                hide-mode-switch
                mode="hexa"
                swatches-max-height="200"
              ></v-color-picker>
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
          <v-btn color="blue darken-1" text @click="dialog = false">
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
  name: "AdminCalendar",
  components: {
    AppBar,
    confirmDialog
  },
  data: () => ({
    apiURL: "/api/v1/calendar",
    toggle: false,
    dialog: false,
    startmenu: false,
    endmenu: false,
    colormenu: false,
    colorpick: "",
    starttime: 0,
    endtime: 0,
    eventData: [],
    startdate: null,
    enddate: null,
    modal3: false,
    modal2: false,
    color: "",
    button: "",
    startepoch: 0,
    endepoch: 0,
    valid: true,
    name: "",
    description: "",
    link: "",
    focus: "",
    type: "month",
    typeToLabel: {
      month: "Month",
      week: "Week",
      day: "Day",
      "4day": "4 Days"
    },
    selectedEvent: {},
    selectedElement: null,
    selectedOpen: false
  }),
  mounted() {
    this.$refs.calendar.checkChange();
  },
  created() {
    this.fetchData();
  },
  methods: {
    viewDay({ date }) {
      this.focus = date;
      this.type = "day";
    },
    getEventColor(event) {
      return event.color;
    },
    setToday() {
      this.focus = "";
    },
    prev() {
      this.$refs.calendar.prev();
    },
    next() {
      this.$refs.calendar.next();
    },
    showEvent({ nativeEvent, event }) {
      const open = () => {
        this.selectedEvent = event;
        this.selectedElement = nativeEvent.target;
        setTimeout(() => {
          this.selectedOpen = true;
        }, 10);
      };

      if (this.selectedOpen) {
        this.selectedOpen = false;
        setTimeout(open, 10);
      } else {
        open();
      }

      nativeEvent.stopPropagation();
    },
    saveForm() {
      this.dialog = true;
      this.clearform();
      this.button = "Save";
    },
    editForm(value) {
      this.id = value.id;
      this.name = value.name;
      this.description = value.description;
      this.link = value.link;
      this.colorpick = value.color;
      this.button = "Update";
      this.dialog = true;
    },
    clearform() {
      this.name = "";
      this.description = "";
      this.link = "";
      this.starttime = "";
      this.startdate = "";
      this.endtime = "";
      this.enddate = "";
      this.colorpick = "";
    },
    epochtodate(epoch) {
      let d = new Date(epoch);
      let day = d.getFullYear() + "-" + d.getMonth() + "-" + d.getDate();
      let time = d.getHours() + ":" + d.getMinutes() + ":" + d.getSeconds();
      return day, time;
    },
    datetoepoch(date) {
      return new Date(date).getTime();
    },
    fetchData() {
      this.axios
        .get(this.apiURL)
        .then(response => {
          this.eventData = response.data["data"];
          if (this.eventData.length > 0) {
            this.eventData.forEach(element => {
              element["timed"] = true;
            });
          }
        })
        .catch(error => {
          console.log(error);
        });
    },
    saveEntity() {
      if (this.$refs.form.validate()) {
        this.$refs.confirm
          .open("Save", "Are you sure?", { color: "green" })
          .then(confirm => {
            if (confirm) {
              this.startepoch = this.datetoepoch(
                this.startdate + "T" + this.starttime
              );
              this.endepoch = this.datetoepoch(
                this.enddate + "T" + this.endtime
              );
              this.axios
                .post(this.apiURL, {
                  name: this.name,
                  description: this.description,
                  link: this.link,
                  color: this.colorpick,
                  start: this.startepoch,
                  end: this.endepoch
                })
                .then(response => {
                  if (response.data.response_status.response_code == 200) {
                    this.clearform();
                    this.dialog = false;
                    this.fetchData();
                  }
                });
            }
          });
      }
    },
    updateEntity() {
      if (this.$refs.form.validate()) {
        this.$refs.confirm
          .open("Update", "Are you sure?", { color: "blue" })
          .then(confirm => {
            if (confirm) {
              this.startepoch = this.datetoepoch(
                this.startdate + "T" + this.starttime
              );
              this.endepoch = this.datetoepoch(
                this.enddate + "T" + this.endtime
              );
              this.axios
                .put(this.apiURL + "/" + this.id, {
                  name: this.name,
                  description: this.description,
                  link: this.link,
                  color: this.colorpick,
                  start: this.startepoch,
                  end: this.endepoch
                })
                .then(response => {
                  if (response.data.response_status.response_code == 200) {
                    this.clearform();
                    this.dialog = false;
                    this.fetchData();
                  }
                });
            }
          });
      }
    },
    deleteEntity(value) {
      this.$refs.confirm
        .open("Delete", "Are you sure?", { color: "red" })
        .then(confirm => {
          if (confirm) {
            this.axios.delete(this.apiURL + "/" + value.id).then(response => {
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
