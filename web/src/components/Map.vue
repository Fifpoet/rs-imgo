<template>
  <div class="map-container" id="map-container"></div>
</template>

<script>
import {latLongToTileXY} from "../utils/latlng.ts";
import {sendGet} from "../api/movemap.ts";

export default {
  name: "mapView",
  components: {},
  data() {
    return {
      map: null,
      OSMUrl: "http://192.168.132.128:8888/v1/map/{z}/{x}/{y}",
      XChangeUrl: "",
      YChangeUrl: "",
      TileX: 0,
      TileY: 0
    };
  },
  mounted() {
    this.map = this.$utils.map.createMap("map-container");
    // 设施地图视图 中心位置
    this.map.setView([0, 0], 3);
    this.map.on("moveend", function (ev) {
      let mapIns = ev.sourceTarget
      let center = mapIns.getCenter();
      let scale = mapIns.getZoom();
      let xy = latLongToTileXY(center.lat, center.lng, scale) //TODO 私服xy转换
      if(mapIns.TileX - xy[0] !== 0 || mapIns.TileY - xy[1] !== 0) {
        console.log("map移动 更新缓存: ", mapIns.TileX - xy[0], mapIns.TileY - xy[1])
        sendGet("xxx")
      }
    })
    this.$utils.map.createTileLayer(this.map, this.OSMUrl, {maxZoom: 8, minZoom: 2});
  },

};
</script>
<style scoped>
.map-container {
  position: absolute;
  left: 0;
  top: 100px;
  width: 100%;
  height: 100%;
}
</style>