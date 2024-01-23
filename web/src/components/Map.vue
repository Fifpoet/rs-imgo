<template>
  <div class="map-container" id="map-container"></div>
</template>

<script>
import {latLongToTileXY} from "../utils/latlng.ts";
import {CacheUpdate} from "../api/movemap.ts";
import {getUserCode} from "../lstore/user";

export default {
  name: "mapView",
  components: {},
  data() {
    return {
      map: null,
      OSMUrl: "http://192.168.132.128:8888/v1/map/{z}/{x}/{y}",
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
      let xy = latLongToTileXY(center.lat, center.lng, scale) //经纬度转换成Tile值
      // 暂时不知道怎么从事件回调中拿到center状态， 所以状态存到后端， 每次moveend发送请求
      let userCode = getUserCode()
      console.log(`http://192.168.132.128:8888/v1/cache/${scale}/${xy[0]}/${xy[1]}?user=${userCode}`)

      CacheUpdate(`http://192.168.132.128:8888/v1/cache/${scale}/${xy[0]}/${xy[1]}?user=${userCode}`)
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