package monitor

const DashboardJSON = `{
   "annotations":{
      "list":[
         {
            "builtIn":1,
            "datasource":"-- Grafana --",
            "enable":true,
            "hide":true,
            "iconColor":"rgba(0, 211, 255, 1)",
            "name":"Annotations & Alerts",
            "type":"dashboard"
         }
      ]
   },
   "editable":true,
   "gnetId":null,
   "graphTooltip":0,
   "id":25,
   "links":[

   ],
   "panels":[
      {
         "aliasColors":{

         },
         "bars":false,
         "dashLength":10,
         "dashes":false,
         "datasource":"prometheus",
         "fill":1,
         "fillGradient":0,
         "gridPos":{
            "h":9,
            "w":12,
            "x":0,
            "y":0
         },
         "hiddenSeries":false,
         "id":2,
         "legend":{
            "avg":false,
            "current":false,
            "max":false,
            "min":false,
            "show":true,
            "total":false,
            "values":false
         },
         "lines":true,
         "linewidth":1,
         "nullPointMode":"null",
         "options":{
            "dataLinks":[

            ]
         },
         "percentage":false,
         "pointradius":2,
         "points":false,
         "renderer":"flot",
         "seriesOverrides":[

         ],
         "spaceLength":10,
         "stack":false,
         "steppedLine":false,
         "targets":[
            {
               "expr":"mysql_global_status_threads_connected",
               "legendFormat":"",
               "refId":"A"
            }
         ],
         "thresholds":[

         ],
         "timeFrom":null,
         "timeRegions":[

         ],
         "timeShift":null,
         "title":"Panel Title",
         "tooltip":{
            "shared":true,
            "sort":0,
            "value_type":"individual"
         },
         "type":"graph",
         "xaxis":{
            "buckets":null,
            "mode":"time",
            "name":null,
            "show":true,
            "values":[

            ]
         },
         "yaxes":[
            {
               "format":"short",
               "label":null,
               "logBase":1,
               "max":null,
               "min":null,
               "show":true
            },
            {
               "format":"short",
               "label":null,
               "logBase":1,
               "max":null,
               "min":null,
               "show":true
            }
         ],
         "yaxis":{
            "align":false,
            "alignLevel":null
         }
      }
   ],
   "schemaVersion":22,
   "style":"dark",
   "tags":[

   ],
   "templating":{
      "list":[

      ]
   },
   "time":{
      "from":"now-6h",
      "to":"now"
   },
   "timepicker":{
      "refresh_intervals":[
         "5s",
         "10s",
         "30s",
         "1m",
         "5m",
         "15m",
         "30m",
         "1h",
         "2h",
         "1d"
      ]
   },
   "timezone":"",
   "title":"MariaDBDashboard",
   "uid":"IEVGlUgMk",
   "version":2
}
`
