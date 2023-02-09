package source

import (
	"encoding/json"
	"fmt"
)

var countryCodeData = `[
  {
    "Countries and Regions": "Angola",
    "国家或地区": "安哥拉",
    "国际域名缩写": "AO",
    "code": "244",
    "时差": "-7"
  },
  {
    "Countries and Regions": "Afghanistan",
    "国家或地区": "阿富汗",
    "国际域名缩写": "AF",
    "code": "93",
    "时差": "0"
  },
  {
    "Countries and Regions": "Albania",
    "国家或地区": "阿尔巴尼亚",
    "国际域名缩写": "AL",
    "code": "355",
    "时差": "-7"
  },
  {
    "Countries and Regions": "Algeria",
    "国家或地区": "阿尔及利亚",
    "国际域名缩写": "DZ",
    "code": "213",
    "时差": "-8"
  },
  {
    "Countries and Regions": "Andorra",
    "国家或地区": "安道尔共和国",
    "国际域名缩写": "AD",
    "code": "376",
    "时差": "-8"
  },
  {
    "Countries and Regions": "Anguilla",
    "国家或地区": "安圭拉岛",
    "国际域名缩写": "AI",
    "code": "1264",
    "时差": "-12"
  },
  {
    "Countries and Regions": "Antigua and Barbuda",
    "国家或地区": "安提瓜和巴布达",
    "国际域名缩写": "AG",
    "code": "1268",
    "时差": "-12"
  },
  {
    "Countries and Regions": "Argentina",
    "国家或地区": "阿根廷",
    "国际域名缩写": "AR",
    "code": "54",
    "时差": "-11"
  },
  {
    "Countries and Regions": "Armenia",
    "国家或地区": "亚美尼亚",
    "国际域名缩写": "AM",
    "code": "374",
    "时差": "-6"
  },
  {
    "Countries and Regions": "Ascension",
    "国家或地区": "阿森松",
    "国际域名缩写": " ",
    "code": "247",
    "时差": "-8"
  },
  {
    "Countries and Regions": "Australia",
    "国家或地区": "澳大利亚",
    "国际域名缩写": "AU",
    "code": "61",
    "时差": "+2"
  },
  {
    "Countries and Regions": "Austria",
    "国家或地区": "奥地利",
    "国际域名缩写": "AT",
    "code": "43",
    "时差": "-7"
  },
  {
    "Countries and Regions": "Azerbaijan",
    "国家或地区": "阿塞拜疆",
    "国际域名缩写": "AZ",
    "code": "994",
    "时差": "-5"
  },
  {
    "Countries and Regions": "Bahamas",
    "国家或地区": "巴哈马",
    "国际域名缩写": "BS",
    "code": "1242",
    "时差": "-13"
  },
  {
    "Countries and Regions": "Bahrain",
    "国家或地区": "巴林",
    "国际域名缩写": "BH",
    "code": "973",
    "时差": "-5"
  },
  {
    "Countries and Regions": "Bangladesh",
    "国家或地区": "孟加拉国",
    "国际域名缩写": "BD",
    "code": "880",
    "时差": "-2"
  },
  {
    "Countries and Regions": "Barbados",
    "国家或地区": "巴巴多斯",
    "国际域名缩写": "BB",
    "code": "1246",
    "时差": "-12"
  },
  {
    "Countries and Regions": "Belarus",
    "国家或地区": "白俄罗斯",
    "国际域名缩写": "BY",
    "code": "375",
    "时差": "-6"
  },
  {
    "Countries and Regions": "Belgium",
    "国家或地区": "比利时",
    "国际域名缩写": "BE",
    "code": "32",
    "时差": "-7"
  },
  {
    "Countries and Regions": "Belize",
    "国家或地区": "伯利兹",
    "国际域名缩写": "BZ",
    "code": "501",
    "时差": "-14"
  },
  {
    "Countries and Regions": "Benin",
    "国家或地区": "贝宁",
    "国际域名缩写": "BJ",
    "code": "229",
    "时差": "-7"
  },
  {
    "Countries and Regions": "Bermuda Is.",
    "国家或地区": "百慕大群岛",
    "国际域名缩写": "BM",
    "code": "1441",
    "时差": "-12"
  },
  {
    "Countries and Regions": "Bolivia",
    "国家或地区": "玻利维亚",
    "国际域名缩写": "BO",
    "code": "591",
    "时差": "-12"
  },
  {
    "Countries and Regions": "Botswana",
    "国家或地区": "博茨瓦纳",
    "国际域名缩写": "BW",
    "code": "267",
    "时差": "-6"
  },
  {
    "Countries and Regions": "Brazil",
    "国家或地区": "巴西",
    "国际域名缩写": "BR",
    "code": "55",
    "时差": "-11"
  },
  {
    "Countries and Regions": "Brunei",
    "国家或地区": "文莱",
    "国际域名缩写": "BN",
    "code": "673",
    "时差": "0"
  },
  {
    "Countries and Regions": "Bulgaria",
    "国家或地区": "保加利亚",
    "国际域名缩写": "BG",
    "code": "359",
    "时差": "-6"
  },
  {
    "Countries and Regions": "Burkina-faso",
    "国家或地区": "布基纳法索",
    "国际域名缩写": "BF",
    "code": "226",
    "时差": "-8"
  },
  {
    "Countries and Regions": "Burma",
    "国家或地区": "缅甸",
    "国际域名缩写": "MM",
    "code": "95",
    "时差": "-1.3"
  },
  {
    "Countries and Regions": "Burundi",
    "国家或地区": "布隆迪",
    "国际域名缩写": "BI",
    "code": "257",
    "时差": "-6"
  },
  {
    "Countries and Regions": "Cameroon",
    "国家或地区": "喀麦隆",
    "国际域名缩写": "CM",
    "code": "237",
    "时差": "-7"
  },
  {
    "Countries and Regions": "Canada",
    "国家或地区": "加拿大",
    "国际域名缩写": "CA",
    "code": "1",
    "时差": "-13"
  },
  {
    "Countries and Regions": "Cayman Is.",
    "国家或地区": "开曼群岛",
    "国际域名缩写": " ",
    "code": "1345",
    "时差": "-13"
  },
  {
    "Countries and Regions": "Central African Republic",
    "国家或地区": "中非共和国",
    "国际域名缩写": "CF",
    "code": "236",
    "时差": "-7"
  },
  {
    "Countries and Regions": "Chad",
    "国家或地区": "乍得",
    "国际域名缩写": "TD",
    "code": "235",
    "时差": "-7"
  },
  {
    "Countries and Regions": "Chile",
    "国家或地区": "智利",
    "国际域名缩写": "CL",
    "code": "56",
    "时差": "-13"
  },
  {
    "Countries and Regions": "China",
    "国家或地区": "中国",
    "国际域名缩写": "CN",
    "code": "86",
    "时差": "0"
  },
  {
    "Countries and Regions": "Colombia",
    "国家或地区": "哥伦比亚",
    "国际域名缩写": "CO",
    "code": "57",
    "时差": "0"
  },
  {
    "Countries and Regions": "Congo",
    "国家或地区": "刚果",
    "国际域名缩写": "CG",
    "code": "242",
    "时差": "-7"
  },
  {
    "Countries and Regions": "Cook Is.",
    "国家或地区": "库克群岛",
    "国际域名缩写": "CK",
    "code": "682",
    "时差": "-18.3"
  },
  {
    "Countries and Regions": "Costa Rica",
    "国家或地区": "哥斯达黎加",
    "国际域名缩写": "CR",
    "code": "506",
    "时差": "-14"
  },
  {
    "Countries and Regions": "Cuba",
    "国家或地区": "古巴",
    "国际域名缩写": "CU",
    "code": "53",
    "时差": "-13"
  },
  {
    "Countries and Regions": "Cyprus",
    "国家或地区": "塞浦路斯",
    "国际域名缩写": "CY",
    "code": "357",
    "时差": "-6"
  },
  {
    "Countries and Regions": "Czech Republic",
    "国家或地区": "捷克",
    "国际域名缩写": "CZ",
    "code": "420",
    "时差": "-7"
  },
  {
    "Countries and Regions": "Denmark",
    "国家或地区": "丹麦",
    "国际域名缩写": "DK",
    "code": "45",
    "时差": "-7"
  },
  {
    "Countries and Regions": "Djibouti",
    "国家或地区": "吉布提",
    "国际域名缩写": "DJ",
    "code": "253",
    "时差": "-5"
  },
  {
    "Countries and Regions": "Dominica Rep.",
    "国家或地区": "多米尼加共和国",
    "国际域名缩写": "DO",
    "code": "1890",
    "时差": "-13"
  },
  {
    "Countries and Regions": "Ecuador",
    "国家或地区": "厄瓜多尔",
    "国际域名缩写": "EC",
    "code": "593",
    "时差": "-13"
  },
  {
    "Countries and Regions": "Egypt",
    "国家或地区": "埃及",
    "国际域名缩写": "EG",
    "code": "20",
    "时差": "-6"
  },
  {
    "Countries and Regions": "EI Salvador",
    "国家或地区": "萨尔瓦多",
    "国际域名缩写": "SV",
    "code": "503",
    "时差": "-14"
  },
  {
    "Countries and Regions": "Estonia",
    "国家或地区": "爱沙尼亚",
    "国际域名缩写": "EE",
    "code": "372",
    "时差": "-5"
  },
  {
    "Countries and Regions": "Ethiopia",
    "国家或地区": "埃塞俄比亚",
    "国际域名缩写": "ET",
    "code": "251",
    "时差": "-5"
  },
  {
    "Countries and Regions": "Fiji",
    "国家或地区": "斐济",
    "国际域名缩写": "FJ",
    "code": "679",
    "时差": "+4"
  },
  {
    "Countries and Regions": "Finland",
    "国家或地区": "芬兰",
    "国际域名缩写": "FI",
    "code": "358",
    "时差": "-6"
  },
  {
    "Countries and Regions": "France",
    "国家或地区": "法国",
    "国际域名缩写": "FR",
    "code": "33",
    "时差": "-8"
  },
  {
    "Countries and Regions": "French Guiana",
    "国家或地区": "法属圭亚那",
    "国际域名缩写": "GF",
    "code": "594",
    "时差": "-12"
  },
  {
    "Countries and Regions": "Gabon",
    "国家或地区": "加蓬",
    "国际域名缩写": "GA",
    "code": "241",
    "时差": "-7"
  },
  {
    "Countries and Regions": "Gambia",
    "国家或地区": "冈比亚",
    "国际域名缩写": "GM",
    "code": "220",
    "时差": "-8"
  },
  {
    "Countries and Regions": "Georgia",
    "国家或地区": "格鲁吉亚",
    "国际域名缩写": "GE",
    "code": "995",
    "时差": "0"
  },
  {
    "Countries and Regions": "Germany",
    "国家或地区": "德国",
    "国际域名缩写": "DE",
    "code": "49",
    "时差": "-7"
  },
  {
    "Countries and Regions": "Ghana",
    "国家或地区": "加纳",
    "国际域名缩写": "GH",
    "code": "233",
    "时差": "-8"
  },
  {
    "Countries and Regions": "Gibraltar",
    "国家或地区": "直布罗陀",
    "国际域名缩写": "GI",
    "code": "350",
    "时差": "-8"
  },
  {
    "Countries and Regions": "Greece",
    "国家或地区": "希腊",
    "国际域名缩写": "GR",
    "code": "30",
    "时差": "-6"
  },
  {
    "Countries and Regions": "Grenada",
    "国家或地区": "格林纳达",
    "国际域名缩写": "GD",
    "code": "1809",
    "时差": "-14"
  },
  {
    "Countries and Regions": "Guam",
    "国家或地区": "关岛",
    "国际域名缩写": "GU",
    "code": "1671",
    "时差": "+2"
  },
  {
    "Countries and Regions": "Guatemala",
    "国家或地区": "危地马拉",
    "国际域名缩写": "GT",
    "code": "502",
    "时差": "-14"
  },
  {
    "Countries and Regions": "Guinea",
    "国家或地区": "几内亚",
    "国际域名缩写": "GN",
    "code": "224",
    "时差": "-8"
  },
  {
    "Countries and Regions": "Guyana",
    "国家或地区": "圭亚那",
    "国际域名缩写": "GY",
    "code": "592",
    "时差": "-11"
  },
  {
    "Countries and Regions": "Haiti",
    "国家或地区": "海地",
    "国际域名缩写": "HT",
    "code": "509",
    "时差": "-13"
  },
  {
    "Countries and Regions": "Honduras",
    "国家或地区": "洪都拉斯",
    "国际域名缩写": "HN",
    "code": "504",
    "时差": "-14"
  },
  {
    "Countries and Regions": "Hongkong",
    "国家或地区": "香港",
    "国际域名缩写": "HK",
    "code": "852",
    "时差": "0"
  },
  {
    "Countries and Regions": "Hungary",
    "国家或地区": "匈牙利",
    "国际域名缩写": "HU",
    "code": "36",
    "时差": "-7"
  },
  {
    "Countries and Regions": "Iceland",
    "国家或地区": "冰岛",
    "国际域名缩写": "IS",
    "code": "354",
    "时差": "-9"
  },
  {
    "Countries and Regions": "India",
    "国家或地区": "印度",
    "国际域名缩写": "IN",
    "code": "91",
    "时差": "-2.3"
  },
  {
    "Countries and Regions": "Indonesia",
    "国家或地区": "印度尼西亚",
    "国际域名缩写": "ID",
    "code": "62",
    "时差": "-0.3"
  },
  {
    "Countries and Regions": "Iran",
    "国家或地区": "伊朗",
    "国际域名缩写": "IR",
    "code": "98",
    "时差": "-4.3"
  },
  {
    "Countries and Regions": "Iraq",
    "国家或地区": "伊拉克",
    "国际域名缩写": "IQ",
    "code": "964",
    "时差": "-5"
  },
  {
    "Countries and Regions": "Ireland",
    "国家或地区": "爱尔兰",
    "国际域名缩写": "IE",
    "code": "353",
    "时差": "-4.3"
  },
  {
    "Countries and Regions": "Israel",
    "国家或地区": "以色列",
    "国际域名缩写": "IL",
    "code": "972",
    "时差": "-6"
  },
  {
    "Countries and Regions": "Italy",
    "国家或地区": "意大利",
    "国际域名缩写": "IT",
    "code": "39",
    "时差": "-7"
  },
  {
    "Countries and Regions": "Ivory Coast",
    "国家或地区": "科特迪瓦",
    "国际域名缩写": " ",
    "code": "225",
    "时差": "-6"
  },
  {
    "Countries and Regions": "Jamaica",
    "国家或地区": "牙买加",
    "国际域名缩写": "JM",
    "code": "1876",
    "时差": "-12"
  },
  {
    "Countries and Regions": "Japan",
    "国家或地区": "日本",
    "国际域名缩写": "JP",
    "code": "81",
    "时差": "+1"
  },
  {
    "Countries and Regions": "Jordan",
    "国家或地区": "约旦",
    "国际域名缩写": "JO",
    "code": "962",
    "时差": "-6"
  },
  {
    "Countries and Regions": "Kampuchea (Cambodia )",
    "国家或地区": "柬埔寨",
    "国际域名缩写": "KH",
    "code": "855",
    "时差": "-1"
  },
  {
    "Countries and Regions": "Kazakstan",
    "国家或地区": "哈萨克斯坦",
    "国际域名缩写": "KZ",
    "code": "327",
    "时差": "-5"
  },
  {
    "Countries and Regions": "Kenya",
    "国家或地区": "肯尼亚",
    "国际域名缩写": "KE",
    "code": "254",
    "时差": "-5"
  },
  {
    "Countries and Regions": "Korea",
    "国家或地区": "韩国",
    "国际域名缩写": "KR",
    "code": "82",
    "时差": "+1"
  },
  {
    "Countries and Regions": "Kuwait",
    "国家或地区": "科威特",
    "国际域名缩写": "KW",
    "code": "965",
    "时差": "-5"
  },
  {
    "Countries and Regions": "Kyrgyzstan",
    "国家或地区": "吉尔吉斯坦",
    "国际域名缩写": "KG",
    "code": "331",
    "时差": "-5"
  },
  {
    "Countries and Regions": "Laos",
    "国家或地区": "老挝",
    "国际域名缩写": "LA",
    "code": "856",
    "时差": "-1"
  },
  {
    "Countries and Regions": "Latvia",
    "国家或地区": "拉脱维亚",
    "国际域名缩写": "LV",
    "code": "371",
    "时差": "-5"
  },
  {
    "Countries and Regions": "Lebanon",
    "国家或地区": "黎巴嫩",
    "国际域名缩写": "LB",
    "code": "961",
    "时差": "-6"
  },
  {
    "Countries and Regions": "Lesotho",
    "国家或地区": "莱索托",
    "国际域名缩写": "LS",
    "code": "266",
    "时差": "-6"
  },
  {
    "Countries and Regions": "Liberia",
    "国家或地区": "利比里亚",
    "国际域名缩写": "LR",
    "code": "231",
    "时差": "-8"
  },
  {
    "Countries and Regions": "Libya",
    "国家或地区": "利比亚",
    "国际域名缩写": "LY",
    "code": "218",
    "时差": "-6"
  },
  {
    "Countries and Regions": "Liechtenstein",
    "国家或地区": "列支敦士登",
    "国际域名缩写": "LI",
    "code": "423",
    "时差": "-7"
  },
  {
    "Countries and Regions": "Lithuania",
    "国家或地区": "立陶宛",
    "国际域名缩写": "LT",
    "code": "370",
    "时差": "-5"
  },
  {
    "Countries and Regions": "Luxembourg",
    "国家或地区": "卢森堡",
    "国际域名缩写": "LU",
    "code": "352",
    "时差": "-7"
  },
  {
    "Countries and Regions": "Macao",
    "国家或地区": "澳门",
    "国际域名缩写": "MO",
    "code": "853",
    "时差": "0"
  },
  {
    "Countries and Regions": "Madagascar",
    "国家或地区": "马达加斯加",
    "国际域名缩写": "MG",
    "code": "261",
    "时差": "-5"
  },
  {
    "Countries and Regions": "Malawi",
    "国家或地区": "马拉维",
    "国际域名缩写": "MW",
    "code": "265",
    "时差": "-6"
  },
  {
    "Countries and Regions": "Malaysia",
    "国家或地区": "马来西亚",
    "国际域名缩写": "MY",
    "code": "60",
    "时差": "-0.5"
  },
  {
    "Countries and Regions": "Maldives",
    "国家或地区": "马尔代夫",
    "国际域名缩写": "MV",
    "code": "960",
    "时差": "-7"
  },
  {
    "Countries and Regions": "Mali",
    "国家或地区": "马里",
    "国际域名缩写": "ML",
    "code": "223",
    "时差": "-8"
  },
  {
    "Countries and Regions": "Malta",
    "国家或地区": "马耳他",
    "国际域名缩写": "MT",
    "code": "356",
    "时差": "-7"
  },
  {
    "Countries and Regions": "Mariana Is",
    "国家或地区": "马里亚那群岛",
    "国际域名缩写": " ",
    "code": "1670",
    "时差": "+1"
  },
  {
    "Countries and Regions": "Martinique",
    "国家或地区": "马提尼克",
    "国际域名缩写": " ",
    "code": "596",
    "时差": "-12"
  },
  {
    "Countries and Regions": "Mauritius",
    "国家或地区": "毛里求斯",
    "国际域名缩写": "MU",
    "code": "230",
    "时差": "-4"
  },
  {
    "Countries and Regions": "Mexico",
    "国家或地区": "墨西哥",
    "国际域名缩写": "MX",
    "code": "52",
    "时差": "-15"
  },
  {
    "Countries and Regions": "Moldova Republic of",
    "国家或地区": "摩尔多瓦",
    "国际域名缩写": "MD",
    "code": "373",
    "时差": "-5"
  },
  {
    "Countries and Regions": "Monaco",
    "国家或地区": "摩纳哥",
    "国际域名缩写": "MC",
    "code": "377",
    "时差": "-7"
  },
  {
    "Countries and Regions": "Mongolia",
    "国家或地区": "蒙古",
    "国际域名缩写": "MN",
    "code": "976",
    "时差": "0"
  },
  {
    "Countries and Regions": "Montserrat Is",
    "国家或地区": "蒙特塞拉特岛",
    "国际域名缩写": "MS",
    "code": "1664",
    "时差": "-12"
  },
  {
    "Countries and Regions": "Morocco",
    "国家或地区": "摩洛哥",
    "国际域名缩写": "MA",
    "code": "212",
    "时差": "-6"
  },
  {
    "Countries and Regions": "Mozambique",
    "国家或地区": "莫桑比克",
    "国际域名缩写": "MZ",
    "code": "258",
    "时差": "-6"
  },
  {
    "Countries and Regions": "Namibia",
    "国家或地区": "纳米比亚",
    "国际域名缩写": "NA",
    "code": "264",
    "时差": "-7"
  },
  {
    "Countries and Regions": "Nauru",
    "国家或地区": "瑙鲁",
    "国际域名缩写": "NR",
    "code": "674",
    "时差": "+4"
  },
  {
    "Countries and Regions": "Nepal",
    "国家或地区": "尼泊尔",
    "国际域名缩写": "NP",
    "code": "977",
    "时差": "-2.3"
  },
  {
    "Countries and Regions": "Netheriands Antilles",
    "国家或地区": "荷属安的列斯",
    "国际域名缩写": " ",
    "code": "599",
    "时差": "-12"
  },
  {
    "Countries and Regions": "Netherlands",
    "国家或地区": "荷兰",
    "国际域名缩写": "NL",
    "code": "31",
    "时差": "-7"
  },
  {
    "Countries and Regions": "New Zealand",
    "国家或地区": "新西兰",
    "国际域名缩写": "NZ",
    "code": "64",
    "时差": "+4"
  },
  {
    "Countries and Regions": "Nicaragua",
    "国家或地区": "尼加拉瓜",
    "国际域名缩写": "NI",
    "code": "505",
    "时差": "-14"
  },
  {
    "Countries and Regions": "Niger",
    "国家或地区": "尼日尔",
    "国际域名缩写": "NE",
    "code": "227",
    "时差": "-8"
  },
  {
    "Countries and Regions": "Nigeria",
    "国家或地区": "尼日利亚",
    "国际域名缩写": "NG",
    "code": "234",
    "时差": "-7"
  },
  {
    "Countries and Regions": "North Korea",
    "国家或地区": "朝鲜",
    "国际域名缩写": "KP",
    "code": "850",
    "时差": "+1"
  },
  {
    "Countries and Regions": "Norway",
    "国家或地区": "挪威",
    "国际域名缩写": "NO",
    "code": "47",
    "时差": "-7"
  },
  {
    "Countries and Regions": "Oman",
    "国家或地区": "阿曼",
    "国际域名缩写": "OM",
    "code": "968",
    "时差": "-4"
  },
  {
    "Countries and Regions": "Pakistan",
    "国家或地区": "巴基斯坦",
    "国际域名缩写": "PK",
    "code": "92",
    "时差": "-2.3"
  },
  {
    "Countries and Regions": "Panama",
    "国家或地区": "巴拿马",
    "国际域名缩写": "PA",
    "code": "507",
    "时差": "-13"
  },
  {
    "Countries and Regions": "Papua New Cuinea",
    "国家或地区": "巴布亚新几内亚",
    "国际域名缩写": "PG",
    "code": "675",
    "时差": "+2"
  },
  {
    "Countries and Regions": "Paraguay",
    "国家或地区": "巴拉圭",
    "国际域名缩写": "PY",
    "code": "595",
    "时差": "-12"
  },
  {
    "Countries and Regions": "Peru",
    "国家或地区": "秘鲁",
    "国际域名缩写": "PE",
    "code": "51",
    "时差": "-13"
  },
  {
    "Countries and Regions": "Philippines",
    "国家或地区": "菲律宾",
    "国际域名缩写": "PH",
    "code": "63",
    "时差": "0"
  },
  {
    "Countries and Regions": "Poland",
    "国家或地区": "波兰",
    "国际域名缩写": "PL",
    "code": "48",
    "时差": "-7"
  },
  {
    "Countries and Regions": "French Polynesia",
    "国家或地区": "法属玻利尼西亚",
    "国际域名缩写": "PF",
    "code": "689",
    "时差": "+3"
  },
  {
    "Countries and Regions": "Portugal",
    "国家或地区": "葡萄牙",
    "国际域名缩写": "PT",
    "code": "351",
    "时差": "-8"
  },
  {
    "Countries and Regions": "Puerto Rico",
    "国家或地区": "波多黎各",
    "国际域名缩写": "PR",
    "code": "1787",
    "时差": "-12"
  },
  {
    "Countries and Regions": "Qatar",
    "国家或地区": "卡塔尔",
    "国际域名缩写": "QA",
    "code": "974",
    "时差": "-5"
  },
  {
    "Countries and Regions": "Reunion",
    "国家或地区": "留尼旺",
    "国际域名缩写": " ",
    "code": "262",
    "时差": "-4"
  },
  {
    "Countries and Regions": "Romania",
    "国家或地区": "罗马尼亚",
    "国际域名缩写": "RO",
    "code": "40",
    "时差": "-6"
  },
  {
    "Countries and Regions": "Russia",
    "国家或地区": "俄罗斯",
    "国际域名缩写": "RU",
    "code": "7",
    "时差": "-5"
  },
  {
    "Countries and Regions": "Saint Lueia",
    "国家或地区": "圣卢西亚",
    "国际域名缩写": "LC",
    "code": "1758",
    "时差": "-12"
  },
  {
    "Countries and Regions": "Saint Vincent",
    "国家或地区": "圣文森特岛",
    "国际域名缩写": "VC",
    "code": "1784",
    "时差": "-12"
  },
  {
    "Countries and Regions": "Samoa Eastern",
    "国家或地区": "东萨摩亚(美)",
    "国际域名缩写": " ",
    "code": "684",
    "时差": "-19"
  },
  {
    "Countries and Regions": "Samoa Western",
    "国家或地区": "西萨摩亚",
    "国际域名缩写": " ",
    "code": "685",
    "时差": "-19"
  },
  {
    "Countries and Regions": "San Marino",
    "国家或地区": "圣马力诺",
    "国际域名缩写": "SM",
    "code": "378",
    "时差": "-7"
  },
  {
    "Countries and Regions": "Sao Tome and Principe",
    "国家或地区": "圣多美和普林西比",
    "国际域名缩写": "ST",
    "code": "239",
    "时差": "-8"
  },
  {
    "Countries and Regions": "Saudi Arabia",
    "国家或地区": "沙特阿拉伯",
    "国际域名缩写": "SA",
    "code": "966",
    "时差": "-5"
  },
  {
    "Countries and Regions": "Senegal",
    "国家或地区": "塞内加尔",
    "国际域名缩写": "SN",
    "code": "221",
    "时差": "-8"
  },
  {
    "Countries and Regions": "Seychelles",
    "国家或地区": "塞舌尔",
    "国际域名缩写": "SC",
    "code": "248",
    "时差": "-4"
  },
  {
    "Countries and Regions": "Sierra Leone",
    "国家或地区": "塞拉利昂",
    "国际域名缩写": "SL",
    "code": "232",
    "时差": "-8"
  },
  {
    "Countries and Regions": "Singapore",
    "国家或地区": "新加坡",
    "国际域名缩写": "SG",
    "code": "65",
    "时差": "+0.3"
  },
  {
    "Countries and Regions": "Slovakia",
    "国家或地区": "斯洛伐克",
    "国际域名缩写": "SK",
    "code": "421",
    "时差": "-7"
  },
  {
    "Countries and Regions": "Slovenia",
    "国家或地区": "斯洛文尼亚",
    "国际域名缩写": "SI",
    "code": "386",
    "时差": "-7"
  },
  {
    "Countries and Regions": "Solomon Is",
    "国家或地区": "所罗门群岛",
    "国际域名缩写": "SB",
    "code": "677",
    "时差": "+3"
  },
  {
    "Countries and Regions": "Somali",
    "国家或地区": "索马里",
    "国际域名缩写": "SO",
    "code": "252",
    "时差": "-5"
  },
  {
    "Countries and Regions": "South Africa",
    "国家或地区": "南非",
    "国际域名缩写": "ZA",
    "code": "27",
    "时差": "-6"
  },
  {
    "Countries and Regions": "Spain",
    "国家或地区": "西班牙",
    "国际域名缩写": "ES",
    "code": "34",
    "时差": "-8"
  },
  {
    "Countries and Regions": "Sri Lanka",
    "国家或地区": "斯里兰卡",
    "国际域名缩写": "LK",
    "code": "94",
    "时差": "0"
  },
  {
    "Countries and Regions": "St.Lucia",
    "国家或地区": "圣卢西亚",
    "国际域名缩写": "LC",
    "code": "1758",
    "时差": "-12"
  },
  {
    "Countries and Regions": "St.Vincent",
    "国家或地区": "圣文森特",
    "国际域名缩写": "VC",
    "code": "1784",
    "时差": "-12"
  },
  {
    "Countries and Regions": "Sudan",
    "国家或地区": "苏丹",
    "国际域名缩写": "SD",
    "code": "249",
    "时差": "-6"
  },
  {
    "Countries and Regions": "Suriname",
    "国家或地区": "苏里南",
    "国际域名缩写": "SR",
    "code": "597",
    "时差": "-11.3"
  },
  {
    "Countries and Regions": "Swaziland",
    "国家或地区": "斯威士兰",
    "国际域名缩写": "SZ",
    "code": "268",
    "时差": "-6"
  },
  {
    "Countries and Regions": "Sweden",
    "国家或地区": "瑞典",
    "国际域名缩写": "SE",
    "code": "46",
    "时差": "-7"
  },
  {
    "Countries and Regions": "Switzerland",
    "国家或地区": "瑞士",
    "国际域名缩写": "CH",
    "code": "41",
    "时差": "-7"
  },
  {
    "Countries and Regions": "Syria",
    "国家或地区": "叙利亚",
    "国际域名缩写": "SY",
    "code": "963",
    "时差": "-6"
  },
  {
    "Countries and Regions": "Taiwan",
    "国家或地区": "台湾省",
    "国际域名缩写": "TW",
    "code": "886",
    "时差": "0"
  },
  {
    "Countries and Regions": "Tajikstan",
    "国家或地区": "塔吉克斯坦",
    "国际域名缩写": "TJ",
    "code": "992",
    "时差": "-5"
  },
  {
    "Countries and Regions": "Tanzania",
    "国家或地区": "坦桑尼亚",
    "国际域名缩写": "TZ",
    "code": "255",
    "时差": "-5"
  },
  {
    "Countries and Regions": "Thailand",
    "国家或地区": "泰国",
    "国际域名缩写": "TH",
    "code": "66",
    "时差": "-1"
  },
  {
    "Countries and Regions": "Togo",
    "国家或地区": "多哥",
    "国际域名缩写": "TG",
    "code": "228",
    "时差": "-8"
  },
  {
    "Countries and Regions": "Tonga",
    "国家或地区": "汤加",
    "国际域名缩写": "TO",
    "code": "676",
    "时差": "+4"
  },
  {
    "Countries and Regions": "Trinidad and Tobago",
    "国家或地区": "特立尼达和多巴哥",
    "国际域名缩写": "TT",
    "code": "1809",
    "时差": "-12"
  },
  {
    "Countries and Regions": "Tunisia",
    "国家或地区": "突尼斯",
    "国际域名缩写": "TN",
    "code": "216",
    "时差": "-7"
  },
  {
    "Countries and Regions": "Turkey",
    "国家或地区": "土耳其",
    "国际域名缩写": "TR",
    "code": "90",
    "时差": "-6"
  },
  {
    "Countries and Regions": "Turkmenistan",
    "国家或地区": "土库曼斯坦",
    "国际域名缩写": "TM",
    "code": "993",
    "时差": "-5"
  },
  {
    "Countries and Regions": "Uganda",
    "国家或地区": "乌干达",
    "国际域名缩写": "UG",
    "code": "256",
    "时差": "-5"
  },
  {
    "Countries and Regions": "Ukraine",
    "国家或地区": "乌克兰",
    "国际域名缩写": "UA",
    "code": "380",
    "时差": "-5"
  },
  {
    "Countries and Regions": "United Arab Emirates",
    "国家或地区": "阿拉伯联合酋长国",
    "国际域名缩写": "AE",
    "code": "971",
    "时差": "-4"
  },
  {
    "Countries and Regions": "United Kiongdom",
    "国家或地区": "英国",
    "国际域名缩写": "GB",
    "code": "44",
    "时差": "-8"
  },
  {
    "Countries and Regions": "United States of America",
    "国家或地区": "美国",
    "国际域名缩写": "US",
    "code": "1",
    "时差": "-13"
  },
  {
    "Countries and Regions": "Uruguay",
    "国家或地区": "乌拉圭",
    "国际域名缩写": "UY",
    "code": "598",
    "时差": "-10.3"
  },
  {
    "Countries and Regions": "Uzbekistan",
    "国家或地区": "乌兹别克斯坦",
    "国际域名缩写": "UZ",
    "code": "233",
    "时差": "-5"
  },
  {
    "Countries and Regions": "Venezuela",
    "国家或地区": "委内瑞拉",
    "国际域名缩写": "VE",
    "code": "58",
    "时差": "-12.3"
  },
  {
    "Countries and Regions": "Vietnam",
    "国家或地区": "越南",
    "国际域名缩写": "VN",
    "code": "84",
    "时差": "-1"
  },
  {
    "Countries and Regions": "Yemen",
    "国家或地区": "也门",
    "国际域名缩写": "YE",
    "code": "967",
    "时差": "-5"
  },
  {
    "Countries and Regions": "Yugoslavia",
    "国家或地区": "南斯拉夫",
    "国际域名缩写": "YU",
    "code": "381",
    "时差": "-7"
  },
  {
    "Countries and Regions": "Zimbabwe",
    "国家或地区": "津巴布韦",
    "国际域名缩写": "ZW",
    "code": "263",
    "时差": "-6"
  },
  {
    "Countries and Regions": "Zaire",
    "国家或地区": "扎伊尔",
    "国际域名缩写": "ZR",
    "code": "243",
    "时差": "-7"
  },
  {
    "Countries and Regions": "Zambia",
    "国家或地区": "赞比亚",
    "国际域名缩写": "ZM",
    "code": "260",
    "时差": "-6"
  }
]`

type CountryCodeModel []*struct {
	Code string `json:"code"`
}

var CountryCodeData CountryCodeModel

func ParseCountryCode() {
	err := json.Unmarshal([]byte(countryCodeData), &CountryCodeData)
	if err != nil {
		fmt.Errorf("解析手机区号失败:%v", err)
		return
	}
}
