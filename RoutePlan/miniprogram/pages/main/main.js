const { getRandomPlaceHolder } = require("../../utils/util")

// pages/main/main.js
Page({

    /**
     * 页面的初始数据
     */
    data: {
        inputBlocks: [],
        inputBlockCount: 0,
        inputBlocksLength: 0,
        latitude: 0,
        longitude: 0,
        inputValue: [],
    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad(options) {
        var randomPlaceHolder = getRandomPlaceHolder()
        this.data.inputBlocks = [{
            placeHolder: randomPlaceHolder,
            key: this.data.inputBlockCount,
        }]
        this.setInputBlocks(this.data.inputBlocks)
        this.setData({
            inputBlockCount: this.data.inputBlockCount + 1,
        })
        this.setData({
            inputValue: [randomPlaceHolder],
        })

        const that = this
        wx.getLocation({
            type: 'wgs84',
            success (res) {
                that.setData({
                    latitude: res.latitude,
                    longitude: res.longitude,
                })
            }
        })
    },

    onInputValueChanged: function(e) {
        var currentIndex = e.target.dataset.index
        if (e.detail.value == "") {
            this.data.inputValue[currentIndex] = this.data.inputBlocks[currentIndex].placeHolder
            this.setData({
                inputValue: this.data.inputValue
            })
            console.log(this.data.inputValue)
            return
        }

        this.data.inputValue[currentIndex] = e.detail.value
        this.setData({
            inputValue: this.data.inputValue
        })
        console.log(this.data.inputValue)
    },

    onGoButtonTapped: function(e) {
        var app = getApp()
        wx.request({
            url: 'https://www.go-route-plan.top/v1/planRoute',
            data: {
                "gps_info":{
                    "latitude": this.data.latitude,
                    "longitude": this.data.longitude,
                }, 
                "keywords": this.data.inputValue,
            },
            method: "POST",
            header: {
                'content-type': 'application/json' // 默认值
            },
            success: function (res) {
                console.log(res)
                // Handle the response data from the server
                var route_infos = res.data.routeInfos
                var routes_data = []
                for (let i=0; i<route_infos.length; i++){
                    var route_data = {}
                    route_data.tag = app.globalData.tagMap.get(route_infos[i].routeTag)

                    route_data.businesses = []
                    var stopPoints = route_infos[i].stopPoints
                    for (let j=0; j<stopPoints.length; j++) {
                        route_data.businesses.push({
                            name: stopPoints[j].name,
                            rating: stopPoints[j].rating,
                            cost: stopPoints[j].cost,
                            latitude: stopPoints[j].gpsInfo.latitude,
                            longitude: stopPoints[j].gpsInfo.longitude,
                        })
                    }
                    
                    routes_data.push(route_data)
                }
                wx.setStorageSync('routes_data', routes_data)
                wx.navigateTo({
                  url: '/pages/routes/routes',
                })
            },
            fail: function (error) {
                // Handle the request error
                console.log('Request failed:', error);
            }
        });
    },

    onAddButtonTapped: function(e) {
        if (this.data.inputBlocks.length == 5) {
            wx.showToast({
                title: '至多5个输入',
                icon: 'none',
                duration: 2000
              })
            return
        }
        var currentIndex = e.target.dataset.index
        var randomPlaceHolder = getRandomPlaceHolder()
        this.data.inputBlocks.splice(currentIndex+1, 0, {
            placeHolder: randomPlaceHolder,
            key: this.data.inputBlockCount,
        })
        this.setInputBlocks(this.data.inputBlocks)
        this.setData({
            inputBlockCount: this.data.inputBlockCount + 1,
        })

        this.data.inputValue.push(randomPlaceHolder)
        this.setData({
            inputValue: this.data.inputValue,
        })
    },

    onDeleteButtonTapped: function(e){
        var currentIndex = e.target.dataset.index
        this.data.inputBlocks.splice(currentIndex, 1)
        this.setInputBlocks(this.data.inputBlocks)

        this.data.inputValue.splice(currentIndex, 1)
        this.setData({
            inputValue: this.data.inputValue,
        })
    },

    setInputBlocks: function(blocks) {
        this.setData({
            inputBlocks: blocks,
            inputBlocksLength: blocks.length,
        })
    },

    onChooseLocationTapped: function(e) {
        const that = this
        wx.chooseLocation({
            success: function(res) {
                that.setData({
                    latitude: res.latitude,
                    longitude: res.longitude,
                })
            }
        })
    },
})