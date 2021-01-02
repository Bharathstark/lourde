<template>
  <v-main>
    <v-banner v-show="banner.description == 'active'">
      {{ banner.name }}
      <template v-slot:actions>
        <v-btn
          v-show="banner.url != ''"
          :href="banner.url"
          target="_blank"
          text
          color="deep-purple accent-4"
        >
          Go to Link
        </v-btn>
        <v-btn text color="deep-purple accent-4" v-on:click="banner = false">
          close
        </v-btn>
      </template>
    </v-banner>
    <v-app-bar fixed inverted-scroll scroll-threshold="300">
      <v-app-bar-nav-icon
        v-show="$vuetify.breakpoint.xs || $vuetify.breakpoint.sm"
        @click.stop="drawer = !drawer"
      >
      </v-app-bar-nav-icon>
      <v-img class="shrink" src="../assets/logos.png" contain height="50px" />
      <v-toolbar-title
        :style="
          $vuetify.breakpoint.xs || $vuetify.breakpoint.sm ? styled2 : styled3
        "
        class="ml-1"
        >Lourde Matha Convent Matric Higher<br />Secondary
        School</v-toolbar-title
      >
      <v-spacer></v-spacer>
      <v-item-group
        v-show="!($vuetify.breakpoint.xs || $vuetify.breakpoint.sm)"
        v-for="(AppBarItem, i) in AppBarItems"
        :key="i"
      >
        <v-btn text :href="AppBarItem.href">
          {{ AppBarItem.display }}
        </v-btn>
      </v-item-group>
      <v-btn icon @click="$refs.notidialog.open()">
        <v-badge overlap dot>
          <v-icon medium right>
            mdi-bell
          </v-icon>
        </v-badge>
      </v-btn>
    </v-app-bar>
    <section id="welcomeSlider">
      <v-carousel
        show-arrows-on-hover
        cycle
        style="height:100%"
        hide-delimiters
      >
        <v-app-bar
          flat
          color="transparent"
          style="position:absolute;top:20px;z-index:1;"
        >
          <v-app-bar-nav-icon
            v-show="$vuetify.breakpoint.xs || $vuetify.breakpoint.sm"
            @click.stop="drawer = !drawer"
          >
          </v-app-bar-nav-icon>
          <v-menu top v-model="drawer">
            <v-list>
              <v-list-item
                v-for="(AppBarItem, index) in AppBarItems"
                :key="index"
              >
                <v-btn text :href="AppBarItem.href">{{
                  AppBarItem.display
                }}</v-btn>
              </v-list-item>
            </v-list>
          </v-menu>
          <v-img
            class="shrink"
            src="../assets/logos.png"
            contain
            height="80px"
          />
          <v-spacer></v-spacer>
          <v-item-group
            v-show="!($vuetify.breakpoint.xs || $vuetify.breakpoint.sm)"
            v-for="(AppBarItem, i) in AppBarItems"
            :key="i"
          >
            <v-btn :href="AppBarItem.href" text>
              {{ AppBarItem.display }}
            </v-btn>
          </v-item-group>
          <v-btn icon @click="$refs.notidialog.open()">
            <v-badge overlap dot>
              <v-icon medium right>
                mdi-bell
              </v-icon>
            </v-badge>
          </v-btn>
        </v-app-bar>
        <v-carousel-item
          v-for="(item, i) in items"
          :key="i"
          :min-height="'calc(100vh - ' + $vuetify.application.top + 'px)'"
        >
          <v-card>
            <v-img
              id="sliderImage"
              class="align-end"
              :src="item.src"
              :height="'calc(100vh - ' + $vuetify.application.top + 'px)'"
            >
              <v-slide-x-transition>
                <v-row>
                  <v-col cols="12" xs="8" sm="8" md="6">
                    <v-card-text class="text-h5 text-md-h3">
                      {{ item.text }}
                    </v-card-text>
                    <div class="py-8"></div>
                  </v-col>
                </v-row>
              </v-slide-x-transition>
            </v-img>
          </v-card>
        </v-carousel-item>
      </v-carousel>
    </section>
    <notification-card ref="notidialog"></notification-card>
    <section id="ourSchool" data-aos="fade-up" data-aos-duration="1000">
      <OurSchool />
    </section>
    <DialogViewer ref="dialog"></DialogViewer>
    <section id="PrincipalAddress">
      <OurPrincipal />
    </section>
    <section id="Academics">
      <Academics />
    </section>
    <section id="Infrastructures" class=" lighten-3">
      <div class="py-12"></div>
      <v-container class="text-center">
        <h2 class="display-2 font-weight-bold mb-3">Infrastructures</h2>

        <v-responsive class="mx-auto mb-12" width="56">
          <v-divider class="mb-1"></v-divider>
          <v-divider></v-divider>
        </v-responsive>
        <v-row>
          <v-col
            v-for="(infra, i) in Infrastructures"
            :key="i"
            cols="12"
            xs="6"
            sm="6"
            md="4"
          >
            <v-card
              class="py-12 px-4"
              @click="showDialog(infra)"
              color=" lighten-5"
              data-aos="zoom-out-right"
              data-aos-duration="2500"
            >
              <v-avatar size="100" tile>
                <v-img large :src="getImgUrl(infra.iconurl)"></v-img>
              </v-avatar>
              <v-card-title
                class="justify-center font-weight-black text-uppercase"
                v-text="infra.name"
              ></v-card-title>
            </v-card>
          </v-col>
        </v-row>
      </v-container>
    </section>

    <section id="Facilities" class=" lighten-3">
      <div class="py-12"></div>
      <v-container class="text-center">
        <h2 class="display-2 font-weight-bold mb-3">Facilities</h2>
        <v-responsive class="mx-auto mb-12" width="56"> </v-responsive>
        <v-row>
          <v-col
            v-for="item in Facilities"
            :key="item.name"
            cols="12"
            xs="6"
            sm="6"
            md="4"
          >
            <v-card
              class="py-12 px-4"
              @click="showDialog(item)"
              color="lighten-5"
              data-aos="zoom-out-left"
              data-aos-duration="2500"
            >
              <v-avatar size="120">
                <v-img
                  height="120"
                  width="120"
                  :src="getImgUrl(item.iconurl)"
                ></v-img>
              </v-avatar>

              <v-card-title
                class="justify-center  text-subtitle-1 font-weight-black text-uppercase"
                v-text="item.name"
              ></v-card-title>
            </v-card>
          </v-col>
        </v-row>
      </v-container>
    </section>
    <section id="Events">
      <div class="py-6"></div>
      <h2 class="text-md-h3 text-sm-h3  mb-3 text-center ">
        Events
        <v-btn text outlined color="deep-purple" to="/events">Explore</v-btn>
      </h2>
      <v-slide-group center-active show-arrows>
        <v-slide-item v-for="event in eventData" :key="event.id">
          <v-card
            :to="'/event?id=' + event.id"
            class="ma-6"
            :height="
              $vuetify.breakpoint.xs || $vuetify.breakpoint.sm ? 180 : 300
            "
            :width="
              $vuetify.breakpoint.xs || $vuetify.breakpoint.sm ? 150 : 300
            "
          >
            <v-img height="75%" :src="event.thumbnail"> </v-img>
            <v-card-title class="text-lg-h6 text-subtitle-2">
              {{ event.name }}
            </v-card-title>
          </v-card>
        </v-slide-item>
      </v-slide-group>
    </section>
    <section id="Achievements">
      <div class="py-6"></div>
      <h2 class="text-md-h3 text-sm-h3  mb-3 text-center ">
        Achievements
        <v-btn text outlined color="deep-purple" to="/achievements"
          >Explore</v-btn
        >
      </h2>
      <v-slide-group center-active show-arrows>
        <v-slide-item
          v-for="achievement in achievementData"
          :key="achievement.id"
        >
          <v-card
            :to="'/achievement?id=' + achievement.id"
            class="ma-6"
            :height="
              $vuetify.breakpoint.xs || $vuetify.breakpoint.sm ? 180 : 300
            "
            :width="
              $vuetify.breakpoint.xs || $vuetify.breakpoint.sm ? 150 : 300
            "
          >
            <v-img height="75%" :src="achievement.thumbnail"> </v-img>
            <v-card-title class="text-lg-h6 text-subtitle-2">
              {{ achievement.name }}
            </v-card-title>
          </v-card>
        </v-slide-item>
      </v-slide-group>
    </section>
    <section id="Videos">
      <div class="py-6"></div>
      <h2 class="text-md-h3 text-sm-h3  mb-3 text-center ">
        Videos
        <v-btn text to="/videos" outlined color="deep-purple">Explore</v-btn>
      </h2>
      <v-slide-group center-active show-arrows>
        <v-slide-item v-for="video in videoData" :key="video.id">
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
              :src="'https://www.youtube-nocookie.com/embed/' + video.url"
              frameborder="0"
              allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
              allowfullscreen
            ></iframe>
            <v-card-title
              class="ml-3 pa-0 text-lg-h6 text-sm-body-1 text-subtitle-2"
            >
              {{ video.name }}
            </v-card-title>
            <v-card-title class="ml-2 pa-1">
              <v-chip
                v-show="!($vuetify.breakpoint.xs || $vuetify.breakpoint.sm)"
                small
                outlined
                color="deep-purple accent-4"
                >{{
                  new Date(video.date).toISOString().substring(0, 10)
                }}</v-chip
              >
            </v-card-title>
          </v-card>
        </v-slide-item>
      </v-slide-group>
    </section>

    <div class="py-6"></div>
    <section>
      <v-card flat width="100%" dark class="lighten-1 text-center">
        <div class="py-6"></div>
        <v-img class="shrink" src="../assets/logos.png" contain height="80px" />
        <v-card-text class="font-weight text-white"
          >Lourde Matha Convent Matriculation Higher Secondary School
          <br />Dhali road, Bodipatti(PO), <br />Udumalpet - 642154, <br />Tamil
          Nadu, <br />
          04252 - 231288<br />
        </v-card-text>

        <v-card-text>
          <v-btn
            v-for="icon in icons"
            :key="icon"
            :href="icon.link"
            target="_blank"
            class="mx-4"
            icon
          >
            <v-icon size="24px">
              {{ icon.icon }}
            </v-icon>
          </v-btn>
        </v-card-text>

        <v-divider></v-divider>

        <v-card-text>
          Copyright Owned by —
          <strong
            >Lourde Matha Convent Metric Higher Secondary
            School,Udumalpet</strong
          >
        </v-card-text>
      </v-card>
    </section>
  </v-main>
</template>

<script>
import "aos/dist/aos.css";
import AOS from "aos";
export default {
  name: "WelcomeSlider",
  components: {
    NotificationCard: () => import("./NotificationCard.vue"),
    DialogViewer: () => import("./DialogViewer.vue"),
    OurSchool: () => import("./OurSchool.vue"),
    OurPrincipal: () => import("./OurPrincipal.vue"),
    Academics: () => import("./Academics.vue"),
  },

  data: () => ({
    eventURL: "/api/v1/events?module=events",
    achievementURL: "/api/v1/events?module=achievements",
    VideoURL: "/api/v1/notifications?module=videos",
    bannerURL: "/api/v1/notifications?module=banner",
    styled2: "font-family:'Coronet';font-size:19px;line-height: 1.0",
    styled3:
      "font-family:'Coronet';font-size:28px;line-height: 0.9;font-weight: bold;",
    eventData: [],
    achievementData: [],
    videoData: [],
    bannerData: [],
    banner: {},
    dialog: false,
    drawer: false,
    AppBarItems: [
      {
        display: "Home",
        href: "/",
      },
      {
        display: "Academics",
        href: "#Academics",
      },
      {
        display: "Achievements",
        href: "/achievements",
      },
      {
        display: "Events",
        href: "/events",
      },
      {
        display: "Infrastructures",
        href: "#Infrastructures",
      },
      {
        display: "Facilities",
        href: "#Facilities",
      },
      {
        display: "Calendar",
        href: "/calendar",
      },
      {
        display: "Contact Us",
        href: "/contact",
      },
    ],
    items: [
      {
        src: "/img/slider/school.jpg",
        text: "Lourde Matha Convent Matric Higher Secondary School",
      },
      {
        src: "/img/slider/smile.JPG",
        text: "A Reflection of Beauty and Joy ",
      },
      {
        src: "/img/slider/thumbs.jpg",
        text: "Lourde Matha Convent Matric Higher Secondary School",
      },
    ],
    Facilities: [
      {
        name: "Smart learning",
        iconurl: "/icons/smart.png",
        headers: [
          {
            heading: "Smart classes with state of the art techs",
            images: [
              "/facilities/smarttech/NSN_6093-min-compressed.jpg",
              "/facilities/smarttech/NSN_6232-min-compressed.jpg",
              "/facilities/smarttech/NSN_6234-min-compressed.jpg",
              "/facilities/smarttech/STIL-Lab.jpg",
              "/facilities/smarttech/TEC-Teaching.jpg",
            ],
          },
        ],
      },
      {
        name: "Expert sessions",
        iconurl: "/icons/expert.png",
        headers: [
          {
            heading: "Orientation classes by Experts",
            images: [
              "/facilities/expertsession/1 (1).jpg",
              "/facilities/expertsession/1 (2).jpg",
              "/facilities/expertsession/1 (3).jpg",
              "/facilities/expertsession/Sessions-by-Expert_5.jpg",
              "/facilities/expertsession/Sessions-by-Expert_6.jpg",
            ],
          },
        ],
      },
      {
        name: "Counselling",
        iconurl: "/icons/counselling.png",
        headers: [
          {
            heading: "Students Councelling",
            images: [
              "/facilities/counselling/Councling_1.jpg",
              "/facilities/counselling/Councling_2.jpg",
              "/facilities/counselling/Councling_3.jpg",
            ],
          },
        ],
      },
      {
        name: "Sports",
        iconurl: "/icons/sports.png",
        headers: [
          {
            heading: "Physical Health and Fitness",
            images: [
              "/facilities/sports/NSN_6319-min-compressed.jpg",
              "/facilities/sports/NSN_6331-min-compressed.jpg",
              "/facilities/sports/NSN_6394-min-compressed.jpg",
              "/facilities/sports/NSN_6396-min-compressed.jpg",
            ],
          },
        ],
      },
      {
        name: "Transport",
        iconurl: "/icons/bus.png",
        headers: [
          {
            heading: "Transport Busses",
            images: [
              "/facilities/transport/NSN_6799-min-compressed.jpg",
              "/facilities/transport/NSN_6805-min-compressed.jpg",
              "/facilities/transport/NSN_6816-min-compressed.jpg",
            ],
          },
        ],
      },
      {
        name: "Parking",
        iconurl: "/icons/park.png",
        headers: [
          {
            heading: "Parking area",
            images: [
              "/facilities/parking/NSN_6207-min-compressed.jpg",
              "/facilities/parking/NSN_6224-min-compressed.jpg",
              "/facilities/parking/NSN_6225-min-compressed.jpg",
            ],
          },
        ],
      },
      {
        name: "Vibrant community ",
        iconurl: "/icons/vibrant.png",
        headers: [],
      },
      {
        name: "Parents meeting",
        iconurl: "/icons/parent.png",
        headers: [
          {
            heading: "Parents meeting",
            images: ["/facilities/meetings/Parents-Meeting.jpg"],
          },
        ],
      },
      {
        name: "Seminars",
        iconurl: "/icons/seminar.png",
        headers: [
          {
            heading: "Seminars by Professionals",
            images: [
              "/facilities/seminar/IMG_20190713_113028-min.jpg",
              "/facilities/seminar/IMG_20190713_113145-min.jpg",
              "/facilities/seminar/IMG_20190713_113154-min.jpg",
            ],
          },
        ],
      },
    ],
    Infrastructures: [
      {
        name: "Smart Campus",
        iconurl: "/icons/smartcampus.png",
        headers: [
          {
            heading: "campus view",
            images: [
              "/infrastructure/campus/Infra Structure_CCTV_1.jpg",
              "/infrastructure/campus/Infra Structure_CCTV_2.jpg",
              "/infrastructure/campus/NSN_5939-min-compressed.jpg",
              "/infrastructure/campus/NSN_6808-min-compressed.jpg",
              "/infrastructure/campus/NSN_6830-min-compressed.jpg",
              "/infrastructure/campus/NSN_6848-min-compressed.jpg",
            ],
          },
        ],
      },
      {
        name: "Advanced Lab's",
        iconurl: "/icons/microscope.png",
        headers: [
          {
            heading: "Laboratary",
            images: [
              "/infrastructure/labs/NSN_6004-min-compressed (1).jpg",
              "/infrastructure/labs/NSN_6243-compressed.jpg",
              "/infrastructure/labs/NSN_6284-min-compressed (1).jpg",
              "/infrastructure/labs/NSN_6671-min-compressed (1).jpg",
              "/infrastructure/labs/NSN_6688-min-compressed.jpg",
              "/infrastructure/labs/NSN_6697-min-compressed.jpg",
              "/infrastructure/labs/NSN_6704-min-compressed.jpg",
              "/infrastructure/labs/NSN_6709-min-compressed.jpg",
              "/infrastructure/labs/NSN_6711-min-compressed.jpg",
              "/infrastructure/labs/NSN_6713-min-compressed.jpg",
              "/infrastructure/labs/NSN_6717-min-compressed.jpg",
            ],
          },
        ],
      },
      {
        name: "Academic Library",
        iconurl: "/icons/books.png",
        headers: [
          {
            heading: "Library",
            images: [
              "/infrastructure/library/Lib_3.jpg",
              "/infrastructure/library/NSN_6081-min-compressed.jpg",
              "/infrastructure/library/NSN_6098-min-compressed.jpg",
              "/infrastructure/library/NSN_6101-min-compressed.jpg",
            ],
          },
        ],
      },
      {
        name: "Pleasant feel garden ",
        iconurl: "/icons/garden.png",
        headers: [
          {
            heading: "Garden",
            images: [
              "/infrastructure/garden/NSN_6185-min-compressed.jpg",
              "/infrastructure/garden/NSN_6198-min-compressed.jpg",
              "/infrastructure/garden/garden2.jpg",
            ],
          },
        ],
      },
      {
        name: "Soft play park",
        iconurl: "/icons/soft.png",
        headers: [
          {
            heading: "Play area",
            images: [
              "/infrastructure/park/Copy of DSC_0008-min-compressed.jpg",
              "/infrastructure/park/DSC_0011-min-compressed.jpg",
              "/infrastructure/park/DSC_0016-min-compressed.jpg",
              "/infrastructure/park/DSC_0017-min-compressed.jpg",
              "/infrastructure/park/DSC_0018-min-compressed.jpg",
              "/infrastructure/park/DSC_0021-min-compressed.jpg",
              "/infrastructure/park/DSC_0033-min-compressed.jpg",
              "/infrastructure/park/DSC_0034-min-compressed.jpg",
              "/infrastructure/park/DSC_0038-min-compressed.jpg",
              "/infrastructure/park/NSN_5959-min-compressed.jpg",
              "/infrastructure/park/NSN_5970-min-compressed.jpg",
              "/infrastructure/park/NSN_5985-min-compressed.jpg",
            ],
          },
        ],
      },
      {
        name: "Spacious conference hall's",
        iconurl: "/icons/hall.png",
        headers: [
          {
            heading: "conference hall",
            images: [
              "/infrastructure/conference/1 (2).jpg",
              "/infrastructure/conference/conference.jpg",
            ],
          },
        ],
      },
    ],
    icons: [
      {
        icon: "mdi-youtube",
        link: "https://www.youtube.com/channel/UC3jmIA7i91OC4o4gboaAnXg",
      },
      {
        icon: "mdi-gmail",
        link: "mailto:lmcm.school@yahoo.com",
      },
    ],
  }),
  created() {
    this.fetchCreateData();
    AOS.init();
  },
  watch: {
    bannerData: function(newValue) {
      if (newValue != undefined) {
        this.banner = newValue[0];
      }
    },
  },

  methods: {
    getImgUrl(pic) {
      return require("../assets" + pic);
    },
    showDialog(data) {
      this.$refs.dialog.open(data);
    },
    async fetchData(url) {
      return this.axios
        .get(url)
        .then((response) => {
          return response.data["data"];
        })
        .catch((error) => {
          console.log(error);
        });
    },
    async fetchCreateData() {
      this.fetchData(this.eventURL).then((data) => (this.eventData = data));

      this.fetchData(this.achievementURL).then(
        (data) => (this.achievementData = data)
      );
      this.fetchData(this.VideoURL).then((data) => (this.videoData = data));
      this.fetchData(this.bannerURL).then((data) => (this.bannerData = data));
    },
  },
};
</script>

<style scoped>
html {
  overflow: hidden;
}
#welcomeSlider {
  height: 100vh;
}
</style>
