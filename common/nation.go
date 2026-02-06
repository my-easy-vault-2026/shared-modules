package common

type NationCode string // 國碼
const (
	NATION_CODE_USA NationCode = "USA"
	NATION_CODE_JPN NationCode = "JPN"
	NATION_CODE_AUT NationCode = "AUT" // EU Austria
	NATION_CODE_BEL NationCode = "BEL" // EU Belgium
	NATION_CODE_BGR NationCode = "BGR" // EU Bulgaria
	NATION_CODE_CYP NationCode = "CYP" // EU Cyprus
	NATION_CODE_CZE NationCode = "CZE" // EU Czech Republic
	NATION_CODE_DEU NationCode = "DEU" // EU Germany
	NATION_CODE_DNK NationCode = "DNK" // EU Denmark
	NATION_CODE_ESP NationCode = "ESP" // EU Spain
	NATION_CODE_EST NationCode = "EST" // EU Estonia
	NATION_CODE_FIN NationCode = "FIN" // EU Finland
	NATION_CODE_FRA NationCode = "FRA" // EU France
	NATION_CODE_GRC NationCode = "GRC" // EU Greece
	NATION_CODE_HUN NationCode = "HUN" // EU Hungary
	NATION_CODE_IRL NationCode = "IRL" // EU Ireland
	NATION_CODE_ITA NationCode = "ITA" // EU Italy
	NATION_CODE_LAT NationCode = "LAT" // EU Latvia
	NATION_CODE_LIT NationCode = "LIT" // EU Lithuania
	NATION_CODE_LUX NationCode = "LUX" // EU Luxembourg
	NATION_CODE_MLT NationCode = "MLT" // EU Malta
	NATION_CODE_NLD NationCode = "NLD" // EU Netherlands
	NATION_CODE_POL NationCode = "POL" // EU Poland
	NATION_CODE_PRT NationCode = "PRT" // EU Portugal
	NATION_CODE_ROU NationCode = "ROU" // EU Romania
	NATION_CODE_SVK NationCode = "SVK" // EU Slovakia
	NATION_CODE_SVN NationCode = "SVN" // EU Slovenia
	NATION_CODE_SWE NationCode = "SWE" // EU Sweden
	NATION_CODE_GBR NationCode = "GBR"
	NATION_CODE_CHN NationCode = "CHN"
	NATION_CODE_CAN NationCode = "CAN"
	NATION_CODE_AUS NationCode = "AUS"
	NATION_CODE_CHE NationCode = "CHE"
	NATION_CODE_HKG NationCode = "HKG"
	NATION_CODE_NZL NationCode = "NZL"
	NATION_CODE_MEX NationCode = "MEX"
	NATION_CODE_SGP NationCode = "SGP"
	NATION_CODE_NOR NationCode = "NOR"
	NATION_CODE_KOR NationCode = "KOR"
	NATION_CODE_TUR NationCode = "TUR"
	NATION_CODE_IND NationCode = "IND"
	NATION_CODE_BRA NationCode = "BRA"
	NATION_CODE_ZAF NationCode = "ZAF"
	NATION_CODE_RUS NationCode = "RUS"
	NATION_CODE_TWN NationCode = "TWN"
	NATION_CODE_ISR NationCode = "ISR"
	NATION_CODE_ARE NationCode = "ARE"
	NATION_CODE_ARG NationCode = "ARG"
	NATION_CODE_MYS NationCode = "MYS"
	NATION_CODE_THA NationCode = "THA"
	NATION_CODE_SAU NationCode = "SAU"
	NATION_CODE_IDN NationCode = "IDN" // 印尼 (Indonesia)
	NATION_CODE_PAK NationCode = "PAK" // 巴基斯坦 (Pakistan)
	NATION_CODE_NGA NationCode = "NGA" // 奈及利亞 (Nigeria)
	NATION_CODE_EGY NationCode = "EGY" // 埃及 (Egypt)
	NATION_CODE_VNM NationCode = "VNM" // 越南 (Vietnam)
	NATION_CODE_PHL NationCode = "PHL" // 菲律賓 (Philippines)
	NATION_CODE_COL NationCode = "COL" // 哥倫比亞 (Colombia)
	NATION_CODE_CHL NationCode = "CHL" // 智利 (Chile)
	NATION_CODE_PER NationCode = "PER" // 秘魯 (Peru)
	NATION_CODE_DZA NationCode = "DZA" // 阿爾及利亞 (Algeria)
	NATION_CODE_MAR NationCode = "MAR" // 摩洛哥 (Morocco)
	NATION_CODE_KEN NationCode = "KEN" // 肯亞 (Kenya)
	NATION_CODE_UKR NationCode = "UKR" // 烏克蘭 (Ukraine)
	NATION_CODE_AFG NationCode = "AFG" // 阿富汗 (Afghanistan)
	NATION_CODE_ETH NationCode = "ETH" // 埃塞俄比亞 (Ethiopia)
	NATION_CODE_COD NationCode = "COD" // 剛果民主共和國 (Democratic Republic of the Congo)
	NATION_CODE_TZA NationCode = "TZA" // 坦尚尼亞 (Tanzania)
	NATION_CODE_UZB NationCode = "UZB" // 烏茲別克 (Uzbekistan)
	NATION_CODE_ALA NationCode = "ALA" // Åland Islands
	NATION_CODE_ALB NationCode = "ALB" // Albania
	NATION_CODE_ASM NationCode = "ASM" // American Samoa
	NATION_CODE_AND NationCode = "AND" // Andorra
	NATION_CODE_AGO NationCode = "AGO" // Angola
	NATION_CODE_AIA NationCode = "AIA" // Anguilla
	NATION_CODE_ATA NationCode = "ATA" // Antarctica
	NATION_CODE_ATG NationCode = "ATG" // Antigua and Barbuda
	NATION_CODE_ARM NationCode = "ARM" // Armenia
	NATION_CODE_ABW NationCode = "ABW" // Aruba
	NATION_CODE_AZE NationCode = "AZE" // Azerbaijan
	NATION_CODE_BHS NationCode = "BHS" // Bahamas
	NATION_CODE_BHR NationCode = "BHR" // Bahrain
	NATION_CODE_BGD NationCode = "BGD" // Bangladesh
	NATION_CODE_BRB NationCode = "BRB" // Barbados
	NATION_CODE_BLR NationCode = "BLR" // Belarus
	NATION_CODE_BLZ NationCode = "BLZ" // Belize
	NATION_CODE_BEN NationCode = "BEN" // Benin
	NATION_CODE_BMU NationCode = "BMU" // Bermuda
	NATION_CODE_BTN NationCode = "BTN" // Bhutan
	NATION_CODE_BOL NationCode = "BOL" // Bolivia (Plurinational State of)
	NATION_CODE_BES NationCode = "BES" // Bonaire, Sint Eustatius and Saba
	NATION_CODE_BIH NationCode = "BIH" // Bosnia and Herzegovina
	NATION_CODE_BWA NationCode = "BWA" // Botswana
	NATION_CODE_BVT NationCode = "BVT" // Bouvet Island
	NATION_CODE_IOT NationCode = "IOT" // British Indian Ocean Territory
	NATION_CODE_BRN NationCode = "BRN" // Brunei Darussalam
	NATION_CODE_BFA NationCode = "BFA" // Burkina Faso
	NATION_CODE_BDI NationCode = "BDI" // Burundi
	NATION_CODE_CPV NationCode = "CPV" // Cabo Verde
	NATION_CODE_KHM NationCode = "KHM" // Cambodia
	NATION_CODE_CMR NationCode = "CMR" // Cameroon
	NATION_CODE_CYM NationCode = "CYM" // Cayman Islands
	NATION_CODE_CAF NationCode = "CAF" // Central African Republic
	NATION_CODE_TCD NationCode = "TCD" // Chad
	NATION_CODE_CXR NationCode = "CXR" // Christmas Island
	NATION_CODE_CCK NationCode = "CCK" // Cocos (Keeling) Islands
	NATION_CODE_COM NationCode = "COM" // Comoros
	NATION_CODE_COG NationCode = "COG" // Congo
	NATION_CODE_COK NationCode = "COK" // Cook Islands
	NATION_CODE_CRI NationCode = "CRI" // Costa Rica
	NATION_CODE_CIV NationCode = "CIV" // Côte d'Ivoire
	NATION_CODE_HRV NationCode = "HRV" // Croatia
	NATION_CODE_CUB NationCode = "CUB" // Cuba
	NATION_CODE_CUW NationCode = "CUW" // Curaçao
	NATION_CODE_DJI NationCode = "DJI" // Djibouti
	NATION_CODE_DMA NationCode = "DMA" // Dominica
	NATION_CODE_DOM NationCode = "DOM" // Dominican Republic
	NATION_CODE_ECU NationCode = "ECU" // Ecuador
	NATION_CODE_SLV NationCode = "SLV" // El Salvador
	NATION_CODE_GNQ NationCode = "GNQ" // Equatorial Guinea
	NATION_CODE_ERI NationCode = "ERI" // Eritrea
	NATION_CODE_SWZ NationCode = "SWZ" // Eswatini
	NATION_CODE_FLK NationCode = "FLK" // Falkland Islands (Malvinas)
	NATION_CODE_FRO NationCode = "FRO" // Faroe Islands
	NATION_CODE_FJI NationCode = "FJI" // Fiji
	NATION_CODE_GUF NationCode = "GUF" // French Guiana
	NATION_CODE_PYF NationCode = "PYF" // French Polynesia
	NATION_CODE_ATF NationCode = "ATF" // French Southern Territories
	NATION_CODE_GAB NationCode = "GAB" // Gabon
	NATION_CODE_GMB NationCode = "GMB" // Gambia
	NATION_CODE_GEO NationCode = "GEO" // Georgia
	NATION_CODE_GHA NationCode = "GHA" // Ghana
	NATION_CODE_GIB NationCode = "GIB" // Gibraltar
	NATION_CODE_GRL NationCode = "GRL" // Greenland
	NATION_CODE_GRD NationCode = "GRD" // Grenada
	NATION_CODE_GLP NationCode = "GLP" // Guadeloupe
	NATION_CODE_GUM NationCode = "GUM" // Guam
	NATION_CODE_GTM NationCode = "GTM" // Guatemala
	NATION_CODE_GGY NationCode = "GGY" // Guernsey
	NATION_CODE_GIN NationCode = "GIN" // Guinea
	NATION_CODE_GNB NationCode = "GNB" // Guinea-Bissau
	NATION_CODE_GUY NationCode = "GUY" // Guyana
	NATION_CODE_HTI NationCode = "HTI" // Haiti
	NATION_CODE_HMD NationCode = "HMD" // Heard Island and McDonald Islands
	NATION_CODE_VAT NationCode = "VAT" // Holy See
	NATION_CODE_HND NationCode = "HND" // Honduras
	NATION_CODE_ISL NationCode = "ISL" // Iceland
	NATION_CODE_IRN NationCode = "IRN" // Iran (Islamic Republic of)
	NATION_CODE_IRQ NationCode = "IRQ" // Iraq
	NATION_CODE_IMN NationCode = "IMN" // Isle of Man
	NATION_CODE_JAM NationCode = "JAM" // Jamaica
	NATION_CODE_JEY NationCode = "JEY" // Jersey
	NATION_CODE_JOR NationCode = "JOR" // Jordan
	NATION_CODE_KAZ NationCode = "KAZ" // Kazakhstan
	NATION_CODE_KIR NationCode = "KIR" // Kiribati
	NATION_CODE_PRK NationCode = "PRK" // Korea (Democratic People's Republic of)
	NATION_CODE_KWT NationCode = "KWT" // Kuwait
	NATION_CODE_KGZ NationCode = "KGZ" // Kyrgyzstan
	NATION_CODE_LAO NationCode = "LAO" // Lao People's Democratic Republic
	NATION_CODE_LVA NationCode = "LVA" // Latvia
	NATION_CODE_LBN NationCode = "LBN" // Lebanon
	NATION_CODE_LSO NationCode = "LSO" // Lesotho
	NATION_CODE_LBR NationCode = "LBR" // Liberia
	NATION_CODE_LBY NationCode = "LBY" // Libya
	NATION_CODE_LIE NationCode = "LIE" // Liechtenstein
	NATION_CODE_LTU NationCode = "LTU" // Lithuania
	NATION_CODE_MAC NationCode = "MAC" // Macao
	NATION_CODE_MDG NationCode = "MDG" // Madagascar
	NATION_CODE_MWI NationCode = "MWI" // Malawi
	NATION_CODE_MDV NationCode = "MDV" // Maldives
	NATION_CODE_MLI NationCode = "MLI" // Mali
	NATION_CODE_MHL NationCode = "MHL" // Marshall Islands
	NATION_CODE_MTQ NationCode = "MTQ" // Martinique
	NATION_CODE_MRT NationCode = "MRT" // Mauritania
	NATION_CODE_MUS NationCode = "MUS" // Mauritius
	NATION_CODE_MYT NationCode = "MYT" // Mayotte
	NATION_CODE_FSM NationCode = "FSM" // Micronesia (Federated States of)
	NATION_CODE_MDA NationCode = "MDA" // Moldova, Republic of
	NATION_CODE_MCO NationCode = "MCO" // Monaco
	NATION_CODE_MNG NationCode = "MNG" // Mongolia
	NATION_CODE_MNE NationCode = "MNE" // Montenegro
	NATION_CODE_MSR NationCode = "MSR" // Montserrat
	NATION_CODE_MOZ NationCode = "MOZ" // Mozambique
	NATION_CODE_MMR NationCode = "MMR" // Myanmar
	NATION_CODE_NAM NationCode = "NAM" // Namibia
	NATION_CODE_NRU NationCode = "NRU" // Nauru
	NATION_CODE_NPL NationCode = "NPL" // Nepal
	NATION_CODE_NCL NationCode = "NCL" // New Caledonia
	NATION_CODE_NIC NationCode = "NIC" // Nicaragua
	NATION_CODE_NER NationCode = "NER" // Niger
	NATION_CODE_NIU NationCode = "NIU" // Niue
	NATION_CODE_NFK NationCode = "NFK" // Norfolk Island
	NATION_CODE_MKD NationCode = "MKD" // North Macedonia
	NATION_CODE_MNP NationCode = "MNP" // Northern Mariana Islands
	NATION_CODE_OMN NationCode = "OMN" // Oman
	NATION_CODE_PLW NationCode = "PLW" // Palau
	NATION_CODE_PSE NationCode = "PSE" // Palestine, State of
	NATION_CODE_PAN NationCode = "PAN" // Panama
	NATION_CODE_PNG NationCode = "PNG" // Papua New Guinea
	NATION_CODE_PRY NationCode = "PRY" // Paraguay
	NATION_CODE_PCN NationCode = "PCN" // Pitcairn
	NATION_CODE_PRI NationCode = "PRI" // Puerto Rico
	NATION_CODE_QAT NationCode = "QAT" // Qatar
	NATION_CODE_REU NationCode = "REU" // Réunion
	NATION_CODE_RWA NationCode = "RWA" // Rwanda
	NATION_CODE_BLM NationCode = "BLM" // Saint Barthélemy
	NATION_CODE_SHN NationCode = "SHN" // Saint Helena, Ascension and Tristan da Cunha
	NATION_CODE_KNA NationCode = "KNA" // Saint Kitts and Nevis
	NATION_CODE_LCA NationCode = "LCA" // Saint Lucia
	NATION_CODE_MAF NationCode = "MAF" // Saint Martin (French part)
	NATION_CODE_SPM NationCode = "SPM" // Saint Pierre and Miquelon
	NATION_CODE_VCT NationCode = "VCT" // Saint Vincent and the Grenadines
	NATION_CODE_WSM NationCode = "WSM" // Samoa
	NATION_CODE_SMR NationCode = "SMR" // San Marino
	NATION_CODE_STP NationCode = "STP" // Sao Tome and Principe
	NATION_CODE_SEN NationCode = "SEN" // Senegal
	NATION_CODE_SRB NationCode = "SRB" // Serbia
	NATION_CODE_SYC NationCode = "SYC" // Seychelles
	NATION_CODE_SLE NationCode = "SLE" // Sierra Leone
	NATION_CODE_SXM NationCode = "SXM" // Sint Maarten (Dutch part)
	NATION_CODE_SLB NationCode = "SLB" // Solomon Islands
	NATION_CODE_SOM NationCode = "SOM" // Somalia
	NATION_CODE_SGS NationCode = "SGS" // South Georgia and the South Sandwich Islands
	NATION_CODE_SSD NationCode = "SSD" // South Sudan
	NATION_CODE_LKA NationCode = "LKA" // Sri Lanka
	NATION_CODE_SDN NationCode = "SDN" // Sudan
	NATION_CODE_SUR NationCode = "SUR" // Suriname
	NATION_CODE_SJM NationCode = "SJM" // Svalbard and Jan Mayen
	NATION_CODE_SYR NationCode = "SYR" // Syrian Arab Republic
	NATION_CODE_TJK NationCode = "TJK" // Tajikistan
	NATION_CODE_TLS NationCode = "TLS" // Timor-Leste
	NATION_CODE_TGO NationCode = "TGO" // Togo
	NATION_CODE_TKL NationCode = "TKL" // Tokelau
	NATION_CODE_TON NationCode = "TON" // Tonga
	NATION_CODE_TTO NationCode = "TTO" // Trinidad and Tobago
	NATION_CODE_TUN NationCode = "TUN" // Tunisia
	NATION_CODE_TKM NationCode = "TKM" // Turkmenistan
	NATION_CODE_TCA NationCode = "TCA" // Turks and Caicos Islands
	NATION_CODE_TUV NationCode = "TUV" // Tuvalu
	NATION_CODE_UGA NationCode = "UGA" // Uganda
	NATION_CODE_UMI NationCode = "UMI" // United States Minor Outlying Islands
	NATION_CODE_URY NationCode = "URY" // Uruguay
	NATION_CODE_VUT NationCode = "VUT" // Vanuatu
	NATION_CODE_VEN NationCode = "VEN" // Venezuela (Bolivarian Republic of)
	NATION_CODE_VGB NationCode = "VGB" // Virgin Islands (British)
	NATION_CODE_VIR NationCode = "VIR" // Virgin Islands (U.S.)
	NATION_CODE_WLF NationCode = "WLF" // Wallis and Futuna
	NATION_CODE_ESH NationCode = "ESH" // Western Sahara
	NATION_CODE_YEM NationCode = "YEM" // Yemen
	NATION_CODE_ZMB NationCode = "ZMB" // Zambia
	NATION_CODE_ZWE NationCode = "ZWE" // Zimbabwe
)

func (nc NationCode) IsValid() bool {
	switch nc {
	case NATION_CODE_USA,
		NATION_CODE_JPN,
		NATION_CODE_AUT,
		NATION_CODE_BEL,
		NATION_CODE_BGR,
		NATION_CODE_CYP,
		NATION_CODE_CZE,
		NATION_CODE_DEU,
		NATION_CODE_DNK,
		NATION_CODE_ESP,
		NATION_CODE_EST,
		NATION_CODE_FIN,
		NATION_CODE_FRA,
		NATION_CODE_GRC,
		NATION_CODE_HUN,
		NATION_CODE_IRL,
		NATION_CODE_ITA,
		NATION_CODE_LAT,
		NATION_CODE_LIT,
		NATION_CODE_LUX,
		NATION_CODE_MLT,
		NATION_CODE_NLD,
		NATION_CODE_POL,
		NATION_CODE_PRT,
		NATION_CODE_ROU,
		NATION_CODE_SVK,
		NATION_CODE_SVN,
		NATION_CODE_SWE,
		NATION_CODE_GBR,
		NATION_CODE_CHN,
		NATION_CODE_CAN,
		NATION_CODE_AUS,
		NATION_CODE_CHE,
		NATION_CODE_HKG,
		NATION_CODE_NZL,
		NATION_CODE_MEX,
		NATION_CODE_SGP,
		NATION_CODE_NOR,
		NATION_CODE_KOR,
		NATION_CODE_TUR,
		NATION_CODE_IND,
		NATION_CODE_BRA,
		NATION_CODE_ZAF,
		NATION_CODE_RUS,
		NATION_CODE_TWN,
		NATION_CODE_ISR,
		NATION_CODE_ARE,
		NATION_CODE_ARG,
		NATION_CODE_MYS,
		NATION_CODE_THA,
		NATION_CODE_SAU,
		NATION_CODE_IDN,
		NATION_CODE_PAK,
		NATION_CODE_NGA,
		NATION_CODE_EGY,
		NATION_CODE_VNM,
		NATION_CODE_PHL,
		NATION_CODE_COL,
		NATION_CODE_CHL,
		NATION_CODE_PER,
		NATION_CODE_DZA,
		NATION_CODE_MAR,
		NATION_CODE_KEN,
		NATION_CODE_UKR,
		NATION_CODE_AFG,
		NATION_CODE_ETH,
		NATION_CODE_COD,
		NATION_CODE_TZA,
		NATION_CODE_UZB,
		NATION_CODE_ALA,
		NATION_CODE_ALB,
		NATION_CODE_ASM,
		NATION_CODE_AND,
		NATION_CODE_AGO,
		NATION_CODE_AIA,
		NATION_CODE_ATA,
		NATION_CODE_ATG,
		NATION_CODE_ARM,
		NATION_CODE_ABW,
		NATION_CODE_AZE,
		NATION_CODE_BHS,
		NATION_CODE_BHR,
		NATION_CODE_BGD,
		NATION_CODE_BRB,
		NATION_CODE_BLR,
		NATION_CODE_BLZ,
		NATION_CODE_BEN,
		NATION_CODE_BMU,
		NATION_CODE_BTN,
		NATION_CODE_BOL,
		NATION_CODE_BES,
		NATION_CODE_BIH,
		NATION_CODE_BWA,
		NATION_CODE_BVT,
		NATION_CODE_IOT,
		NATION_CODE_BRN,
		NATION_CODE_BFA,
		NATION_CODE_BDI,
		NATION_CODE_CPV,
		NATION_CODE_KHM,
		NATION_CODE_CMR,
		NATION_CODE_CYM,
		NATION_CODE_CAF,
		NATION_CODE_TCD,
		NATION_CODE_CXR,
		NATION_CODE_CCK,
		NATION_CODE_COM,
		NATION_CODE_COG,
		NATION_CODE_COK,
		NATION_CODE_CRI,
		NATION_CODE_CIV,
		NATION_CODE_HRV,
		NATION_CODE_CUB,
		NATION_CODE_CUW,
		NATION_CODE_DJI,
		NATION_CODE_DMA,
		NATION_CODE_DOM,
		NATION_CODE_ECU,
		NATION_CODE_SLV,
		NATION_CODE_GNQ,
		NATION_CODE_ERI,
		NATION_CODE_SWZ,
		NATION_CODE_FLK,
		NATION_CODE_FRO,
		NATION_CODE_FJI,
		NATION_CODE_GUF,
		NATION_CODE_PYF,
		NATION_CODE_ATF,
		NATION_CODE_GAB,
		NATION_CODE_GMB,
		NATION_CODE_GEO,
		NATION_CODE_GHA,
		NATION_CODE_GIB,
		NATION_CODE_GRL,
		NATION_CODE_GRD,
		NATION_CODE_GLP,
		NATION_CODE_GUM,
		NATION_CODE_GTM,
		NATION_CODE_GGY,
		NATION_CODE_GIN,
		NATION_CODE_GNB,
		NATION_CODE_GUY,
		NATION_CODE_HTI,
		NATION_CODE_HMD,
		NATION_CODE_VAT,
		NATION_CODE_HND,
		NATION_CODE_ISL,
		NATION_CODE_IRN,
		NATION_CODE_IRQ,
		NATION_CODE_IMN,
		NATION_CODE_JAM,
		NATION_CODE_JEY,
		NATION_CODE_JOR,
		NATION_CODE_KAZ,
		NATION_CODE_KIR,
		NATION_CODE_PRK,
		NATION_CODE_KWT,
		NATION_CODE_KGZ,
		NATION_CODE_LAO,
		NATION_CODE_LVA,
		NATION_CODE_LBN,
		NATION_CODE_LSO,
		NATION_CODE_LBR,
		NATION_CODE_LBY,
		NATION_CODE_LIE,
		NATION_CODE_LTU,
		NATION_CODE_MAC,
		NATION_CODE_MDG,
		NATION_CODE_MWI,
		NATION_CODE_MDV,
		NATION_CODE_MLI,
		NATION_CODE_MHL,
		NATION_CODE_MTQ,
		NATION_CODE_MRT,
		NATION_CODE_MUS,
		NATION_CODE_MYT,
		NATION_CODE_FSM,
		NATION_CODE_MDA,
		NATION_CODE_MCO,
		NATION_CODE_MNG,
		NATION_CODE_MNE,
		NATION_CODE_MSR,
		NATION_CODE_MOZ,
		NATION_CODE_MMR,
		NATION_CODE_NAM,
		NATION_CODE_NRU,
		NATION_CODE_NPL,
		NATION_CODE_NCL,
		NATION_CODE_NIC,
		NATION_CODE_NER,
		NATION_CODE_NIU,
		NATION_CODE_NFK,
		NATION_CODE_MKD,
		NATION_CODE_MNP,
		NATION_CODE_OMN,
		NATION_CODE_PLW,
		NATION_CODE_PSE,
		NATION_CODE_PAN,
		NATION_CODE_PNG,
		NATION_CODE_PRY,
		NATION_CODE_PCN,
		NATION_CODE_PRI,
		NATION_CODE_QAT,
		NATION_CODE_REU,
		NATION_CODE_RWA,
		NATION_CODE_BLM,
		NATION_CODE_SHN,
		NATION_CODE_KNA,
		NATION_CODE_LCA,
		NATION_CODE_MAF,
		NATION_CODE_SPM,
		NATION_CODE_VCT,
		NATION_CODE_WSM,
		NATION_CODE_SMR,
		NATION_CODE_STP,
		NATION_CODE_SEN,
		NATION_CODE_SRB,
		NATION_CODE_SYC,
		NATION_CODE_SLE,
		NATION_CODE_SXM,
		NATION_CODE_SLB,
		NATION_CODE_SOM,
		NATION_CODE_SGS,
		NATION_CODE_SSD,
		NATION_CODE_LKA,
		NATION_CODE_SDN,
		NATION_CODE_SUR,
		NATION_CODE_SJM,
		NATION_CODE_SYR,
		NATION_CODE_TJK,
		NATION_CODE_TLS,
		NATION_CODE_TGO,
		NATION_CODE_TKL,
		NATION_CODE_TON,
		NATION_CODE_TTO,
		NATION_CODE_TUN,
		NATION_CODE_TKM,
		NATION_CODE_TCA,
		NATION_CODE_TUV,
		NATION_CODE_UGA,
		NATION_CODE_UMI,
		NATION_CODE_URY,
		NATION_CODE_VUT,
		NATION_CODE_VEN,
		NATION_CODE_VGB,
		NATION_CODE_VIR,
		NATION_CODE_WLF,
		NATION_CODE_ESH,
		NATION_CODE_YEM,
		NATION_CODE_ZMB,
		NATION_CODE_ZWE:
		return true
	default:
		return false
	}
}

func (nc NationCode) NationCodeFromISO4217(s string) NationCode {
	switch s {
	case "840":
		return NATION_CODE_USA
	case "392":
		return NATION_CODE_JPN
	case "040":
		return NATION_CODE_AUT
	case "056":
		return NATION_CODE_BEL
	case "100":
		return NATION_CODE_BGR
	case "196":
		return NATION_CODE_CYP
	case "203":
		return NATION_CODE_CZE
	case "276":
		return NATION_CODE_DEU
	case "208":
		return NATION_CODE_DNK
	case "220":
		return NATION_CODE_ESP
	case "233":
		return NATION_CODE_EST
	case "246":
		return NATION_CODE_FIN
	case "250":
		return NATION_CODE_FRA
	case "300":
		return NATION_CODE_GRC
	case "348":
		return NATION_CODE_HUN
	case "372":
		return NATION_CODE_IRL
	case "380":
		return NATION_CODE_ITA
	case "428":
		return NATION_CODE_LAT
	case "440":
		return NATION_CODE_LIT
	case "442":
		return NATION_CODE_LUX
	case "470":
		return NATION_CODE_MLT
	case "528":
		return NATION_CODE_NLD
	case "616":
		return NATION_CODE_POL
	case "620":
		return NATION_CODE_PRT
	case "642":
		return NATION_CODE_ROU
	case "703":
		return NATION_CODE_SVK
	case "705":
		return NATION_CODE_SVN
	case "752":
		return NATION_CODE_SWE
	case "826":
		return NATION_CODE_GBR
	case "156":
		return NATION_CODE_CHN
	case "124":
		return NATION_CODE_CAN
	case "036":
		return NATION_CODE_AUS
	case "756":
		return NATION_CODE_CHE
	case "344":
		return NATION_CODE_HKG
	case "554":
		return NATION_CODE_NZL
	case "484":
		return NATION_CODE_MEX
	case "702":
		return NATION_CODE_SGP
	case "578":
		return NATION_CODE_NOR
	case "410":
		return NATION_CODE_KOR
	case "949":
		return NATION_CODE_TUR
	case "356":
		return NATION_CODE_IND
	case "986":
		return NATION_CODE_BRA
	case "710":
		return NATION_CODE_ZAF
	case "643":
		return NATION_CODE_RUS
	case "158":
		return NATION_CODE_TWN
	case "376":
		return NATION_CODE_ISR
	case "784":
		return NATION_CODE_ARE
	case "032":
		return NATION_CODE_ARG
	case "458":
		return NATION_CODE_MYS
	case "764":
		return NATION_CODE_THA
	case "682":
		return NATION_CODE_SAU
	case "360": // 印度尼西亞 (Indonesia)
		return NATION_CODE_IDN
	case "586": // 巴基斯坦 (Pakistan)
		return NATION_CODE_PAK
	case "566": // 奈及利亞 (Nigeria)
		return NATION_CODE_NGA
	case "818": // 埃及 (Egypt)
		return NATION_CODE_EGY
	case "704": // 越南 (Vietnam)
		return NATION_CODE_VNM
	case "608": // 菲律賓 (Philippines)
		return NATION_CODE_PHL
	case "170": // 哥倫比亞 (Colombia)
		return NATION_CODE_COL
	case "152": // 智利 (Chile)
		return NATION_CODE_CHL
	case "604": // 秘魯 (Peru)
		return NATION_CODE_PER
	case "012": // 阿爾及利亞 (Algeria)
		return NATION_CODE_DZA
	case "504": // 摩洛哥 (Morocco)
		return NATION_CODE_MAR
	case "404": // 肯亞 (Kenya)
		return NATION_CODE_KEN
	case "804": // 烏克蘭 (Ukraine)
		return NATION_CODE_UKR
	case "004": // 阿富汗 (Afghanistan)
		return NATION_CODE_AFG
	case "231": // 埃塞俄比亞 (Ethiopia)
		return NATION_CODE_ETH
	case "180": // 剛果民主共和國 (Democratic Republic of the Congo)
		return NATION_CODE_COD
	case "834": // 坦尚尼亞 (Tanzania)
		return NATION_CODE_TZA
	case "860": // 烏茲別克 (Uzbekistan)
		return NATION_CODE_UZB

	case "248":
		return NATION_CODE_ALA
	case "008":
		return NATION_CODE_ALB
	case "016":
		return NATION_CODE_ASM
	case "020":
		return NATION_CODE_AND
	case "024":
		return NATION_CODE_AGO
	case "660":
		return NATION_CODE_AIA
	case "010":
		return NATION_CODE_ATA
	case "028":
		return NATION_CODE_ATG
	case "051":
		return NATION_CODE_ARM
	case "533":
		return NATION_CODE_ABW
	case "031":
		return NATION_CODE_AZE
	case "044":
		return NATION_CODE_BHS
	case "048":
		return NATION_CODE_BHR
	case "050":
		return NATION_CODE_BGD
	case "052":
		return NATION_CODE_BRB
	case "112":
		return NATION_CODE_BLR
	case "084":
		return NATION_CODE_BLZ
	case "204":
		return NATION_CODE_BEN
	case "060":
		return NATION_CODE_BMU
	case "064":
		return NATION_CODE_BTN
	case "068":
		return NATION_CODE_BOL
	case "535":
		return NATION_CODE_BES
	case "070":
		return NATION_CODE_BIH
	case "072":
		return NATION_CODE_BWA
	case "074":
		return NATION_CODE_BVT
	case "076":
		return NATION_CODE_BRA
	case "086":
		return NATION_CODE_IOT
	case "096":
		return NATION_CODE_BRN
	case "854":
		return NATION_CODE_BFA
	case "108":
		return NATION_CODE_BDI
	case "132":
		return NATION_CODE_CPV
	case "116":
		return NATION_CODE_KHM
	case "120":
		return NATION_CODE_CMR
	case "136":
		return NATION_CODE_CYM
	case "140":
		return NATION_CODE_CAF
	case "148":
		return NATION_CODE_TCD
	case "162":
		return NATION_CODE_CXR
	case "166":
		return NATION_CODE_CCK
	case "174":
		return NATION_CODE_COM
	case "178":
		return NATION_CODE_COG
	case "184":
		return NATION_CODE_COK
	case "188":
		return NATION_CODE_CRI
	case "384":
		return NATION_CODE_CIV
	case "191":
		return NATION_CODE_HRV
	case "192":
		return NATION_CODE_CUB
	case "531":
		return NATION_CODE_CUW
	case "262":
		return NATION_CODE_DJI
	case "212":
		return NATION_CODE_DMA
	case "214":
		return NATION_CODE_DOM
	case "218":
		return NATION_CODE_ECU
	case "222":
		return NATION_CODE_SLV
	case "226":
		return NATION_CODE_GNQ
	case "232":
		return NATION_CODE_ERI
	case "748":
		return NATION_CODE_SWZ
	case "238":
		return NATION_CODE_FLK
	case "234":
		return NATION_CODE_FRO
	case "242":
		return NATION_CODE_FJI
	case "254":
		return NATION_CODE_GUF
	case "258":
		return NATION_CODE_PYF
	case "260":
		return NATION_CODE_ATF
	case "266":
		return NATION_CODE_GAB
	case "270":
		return NATION_CODE_GMB
	case "268":
		return NATION_CODE_GEO
	case "288":
		return NATION_CODE_GHA
	case "292":
		return NATION_CODE_GIB
	case "304":
		return NATION_CODE_GRL
	case "308":
		return NATION_CODE_GRD
	case "312":
		return NATION_CODE_GLP
	case "316":
		return NATION_CODE_GUM
	case "320":
		return NATION_CODE_GTM
	case "831":
		return NATION_CODE_GGY
	case "324":
		return NATION_CODE_GIN
	case "624":
		return NATION_CODE_GNB
	case "328":
		return NATION_CODE_GUY
	case "332":
		return NATION_CODE_HTI
	case "334":
		return NATION_CODE_HMD
	case "336":
		return NATION_CODE_VAT
	case "340":
		return NATION_CODE_HND
	case "352":
		return NATION_CODE_ISL
	case "364":
		return NATION_CODE_IRN
	case "368":
		return NATION_CODE_IRQ
	case "833":
		return NATION_CODE_IMN
	case "388":
		return NATION_CODE_JAM
	case "832":
		return NATION_CODE_JEY
	case "400":
		return NATION_CODE_JOR
	case "398":
		return NATION_CODE_KAZ
	case "296":
		return NATION_CODE_KIR
	case "408":
		return NATION_CODE_PRK
	case "414":
		return NATION_CODE_KWT
	case "417":
		return NATION_CODE_KGZ
	case "418":
		return NATION_CODE_LAO
	case "422":
		return NATION_CODE_LBN
	case "426":
		return NATION_CODE_LSO
	case "430":
		return NATION_CODE_LBR
	case "434":
		return NATION_CODE_LBY
	case "438":
		return NATION_CODE_LIE
	case "446":
		return NATION_CODE_MAC
	case "450":
		return NATION_CODE_MDG
	case "454":
		return NATION_CODE_MWI
	case "462":
		return NATION_CODE_MDV
	case "466":
		return NATION_CODE_MLI
	case "584":
		return NATION_CODE_MHL
	case "474":
		return NATION_CODE_MTQ
	case "478":
		return NATION_CODE_MRT
	case "480":
		return NATION_CODE_MUS
	case "175":
		return NATION_CODE_MYT
	case "583":
		return NATION_CODE_FSM
	case "498":
		return NATION_CODE_MDA
	case "492":
		return NATION_CODE_MCO
	case "496":
		return NATION_CODE_MNG
	case "499":
		return NATION_CODE_MNE
	case "500":
		return NATION_CODE_MSR
	case "508":
		return NATION_CODE_MOZ
	case "104":
		return NATION_CODE_MMR
	case "516":
		return NATION_CODE_NAM
	case "520":
		return NATION_CODE_NRU
	case "524":
		return NATION_CODE_NPL
	case "540":
		return NATION_CODE_NCL
	case "558":
		return NATION_CODE_NIC
	case "562":
		return NATION_CODE_NER
	case "570":
		return NATION_CODE_NIU
	case "574":
		return NATION_CODE_NFK
	case "807":
		return NATION_CODE_MKD
	case "580":
		return NATION_CODE_MNP
	case "512":
		return NATION_CODE_OMN
	case "585":
		return NATION_CODE_PLW
	case "275":
		return NATION_CODE_PSE
	case "591":
		return NATION_CODE_PAN
	case "598":
		return NATION_CODE_PNG
	case "600":
		return NATION_CODE_PRY
	case "612":
		return NATION_CODE_PCN
	case "630":
		return NATION_CODE_PRI
	case "634":
		return NATION_CODE_QAT
	case "638":
		return NATION_CODE_REU
	case "646":
		return NATION_CODE_RWA
	case "652":
		return NATION_CODE_BLM
	case "654":
		return NATION_CODE_SHN
	case "659":
		return NATION_CODE_KNA
	case "662":
		return NATION_CODE_LCA
	case "663":
		return NATION_CODE_MAF
	case "666":
		return NATION_CODE_SPM
	case "670":
		return NATION_CODE_VCT
	case "882":
		return NATION_CODE_WSM
	case "674":
		return NATION_CODE_SMR
	case "678":
		return NATION_CODE_STP
	case "686":
		return NATION_CODE_SEN
	case "688":
		return NATION_CODE_SRB
	case "690":
		return NATION_CODE_SYC
	case "694":
		return NATION_CODE_SLE
	case "534":
		return NATION_CODE_SXM
	case "090":
		return NATION_CODE_SLB
	case "706":
		return NATION_CODE_SOM
	case "239":
		return NATION_CODE_SGS
	case "728":
		return NATION_CODE_SSD
	case "724":
		return NATION_CODE_ESP
	case "144":
		return NATION_CODE_LKA
	case "729":
		return NATION_CODE_SDN
	case "740":
		return NATION_CODE_SUR
	case "744":
		return NATION_CODE_SJM
	case "760":
		return NATION_CODE_SYR
	case "762":
		return NATION_CODE_TJK
	case "626":
		return NATION_CODE_TLS
	case "768":
		return NATION_CODE_TGO
	case "772":
		return NATION_CODE_TKL
	case "776":
		return NATION_CODE_TON
	case "780":
		return NATION_CODE_TTO
	case "788":
		return NATION_CODE_TUN
	case "792":
		return NATION_CODE_TUR
	case "795":
		return NATION_CODE_TKM
	case "796":
		return NATION_CODE_TCA
	case "798":
		return NATION_CODE_TUV
	case "800":
		return NATION_CODE_UGA
	case "581":
		return NATION_CODE_UMI
	case "858":
		return NATION_CODE_URY
	case "548":
		return NATION_CODE_VUT
	case "862":
		return NATION_CODE_VEN
	case "092":
		return NATION_CODE_VGB
	case "850":
		return NATION_CODE_VIR
	case "876":
		return NATION_CODE_WLF
	case "732":
		return NATION_CODE_ESH
	case "887":
		return NATION_CODE_YEM
	case "894":
		return NATION_CODE_ZMB
	case "716":
		return NATION_CODE_ZWE
	default:
		return ""
	}
}

func (nc NationCode) ToTwoLetter() string {
	var converter = map[NationCode]string{
		NATION_CODE_USA: "US",
		NATION_CODE_JPN: "JP",
		NATION_CODE_AUT: "AT",
		NATION_CODE_BEL: "BE",
		NATION_CODE_BGR: "BG",
		NATION_CODE_CYP: "CY",
		NATION_CODE_CZE: "CZ",
		NATION_CODE_DEU: "DE",
		NATION_CODE_DNK: "DK",
		NATION_CODE_ESP: "ES",
		NATION_CODE_EST: "EE",
		NATION_CODE_FIN: "FI",
		NATION_CODE_FRA: "FR",
		NATION_CODE_GRC: "GR",
		NATION_CODE_HUN: "HU",
		NATION_CODE_IRL: "IE",
		NATION_CODE_ITA: "IT",
		NATION_CODE_LAT: "LV",
		NATION_CODE_LIT: "LT",
		NATION_CODE_LUX: "LU",
		NATION_CODE_MLT: "MT",
		NATION_CODE_NLD: "NL",
		NATION_CODE_POL: "PL",
		NATION_CODE_PRT: "PT",
		NATION_CODE_ROU: "RO",
		NATION_CODE_SVK: "SK",
		NATION_CODE_SVN: "SI",
		NATION_CODE_SWE: "SE",
		NATION_CODE_GBR: "GB",
		NATION_CODE_CHN: "CN",
		NATION_CODE_CAN: "CA",
		NATION_CODE_AUS: "AU",
		NATION_CODE_CHE: "CH",
		NATION_CODE_HKG: "HK",
		NATION_CODE_NZL: "NZ",
		NATION_CODE_MEX: "MX",
		NATION_CODE_SGP: "SG",
		NATION_CODE_NOR: "NO",
		NATION_CODE_KOR: "KR",
		NATION_CODE_TUR: "TR",
		NATION_CODE_IND: "IN",
		NATION_CODE_BRA: "BR",
		NATION_CODE_ZAF: "ZA",
		NATION_CODE_RUS: "RU",
		NATION_CODE_TWN: "TW",
		NATION_CODE_ISR: "IL",
		NATION_CODE_ARE: "AE",
		NATION_CODE_ARG: "AR",
		NATION_CODE_MYS: "MY",
		NATION_CODE_THA: "TH",
		NATION_CODE_SAU: "SA",
		NATION_CODE_IDN: "ID",
		NATION_CODE_PAK: "PK",
		NATION_CODE_NGA: "NG",
		NATION_CODE_EGY: "EG",
		NATION_CODE_VNM: "VN",
		NATION_CODE_PHL: "PH",
		NATION_CODE_COL: "CO",
		NATION_CODE_CHL: "CL",
		NATION_CODE_PER: "PE",
		NATION_CODE_DZA: "DZ",
		NATION_CODE_MAR: "MA",
		NATION_CODE_KEN: "KE",
		NATION_CODE_UKR: "UA",
		NATION_CODE_AFG: "AF",
		NATION_CODE_ETH: "ET",
		NATION_CODE_COD: "CD",
		NATION_CODE_TZA: "TZ",
		NATION_CODE_UZB: "UZ",
		NATION_CODE_ALA: "AX", // Åland Islands
		NATION_CODE_ALB: "AL", // Albania
		NATION_CODE_ASM: "AS", // American Samoa
		NATION_CODE_AND: "AD", // Andorra
		NATION_CODE_AGO: "AO", // Angola
		NATION_CODE_AIA: "AI", // Anguilla
		NATION_CODE_ATA: "AQ", // Antarctica
		NATION_CODE_ATG: "AG", // Antigua and Barbuda
		NATION_CODE_ARM: "AM", // Armenia
		NATION_CODE_ABW: "AW", // Aruba
		NATION_CODE_AZE: "AZ", // Azerbaijan
		NATION_CODE_BHS: "BS", // Bahamas
		NATION_CODE_BHR: "BH", // Bahrain
		NATION_CODE_BGD: "BD", // Bangladesh
		NATION_CODE_BRB: "BB", // Barbados
		NATION_CODE_BLR: "BY", // Belarus
		NATION_CODE_BLZ: "BZ", // Belize
		NATION_CODE_BEN: "BJ", // Benin
		NATION_CODE_BMU: "BM", // Bermuda
		NATION_CODE_BTN: "BT", // Bhutan
		NATION_CODE_BOL: "BO", // Bolivia (Plurinational State of)
		NATION_CODE_BES: "BQ", // Bonaire, Sint Eustatius and Saba
		NATION_CODE_BIH: "BA", // Bosnia and Herzegovina
		NATION_CODE_BWA: "BW", // Botswana
		NATION_CODE_BVT: "BV", // Bouvet Island
		NATION_CODE_IOT: "IO", // British Indian Ocean Territory
		NATION_CODE_BRN: "BN", // Brunei Darussalam
		NATION_CODE_BFA: "BF", // Burkina Faso
		NATION_CODE_BDI: "BI", // Burundi
		NATION_CODE_CPV: "CV", // Cabo Verde
		NATION_CODE_KHM: "KH", // Cambodia
		NATION_CODE_CMR: "CM", // Cameroon
		NATION_CODE_CYM: "KY", // Cayman Islands
		NATION_CODE_CAF: "CF", // Central African Republic
		NATION_CODE_TCD: "TD", // Chad
		NATION_CODE_CXR: "CX", // Christmas Island
		NATION_CODE_CCK: "CC", // Cocos (Keeling) Islands
		NATION_CODE_COM: "KM", // Comoros
		NATION_CODE_COG: "CG", // Congo
		NATION_CODE_COK: "CK", // Cook Islands
		NATION_CODE_CRI: "CR", // Costa Rica
		NATION_CODE_CIV: "CI", // Côte d'Ivoire
		NATION_CODE_HRV: "HR", // Croatia
		NATION_CODE_CUB: "CU", // Cuba
		NATION_CODE_CUW: "CW", // Curaçao
		NATION_CODE_DJI: "DJ", // Djibouti
		NATION_CODE_DMA: "DM", // Dominica
		NATION_CODE_DOM: "DO", // Dominican Republic
		NATION_CODE_ECU: "EC", // Ecuador
		NATION_CODE_SLV: "SV", // El Salvador
		NATION_CODE_GNQ: "GQ", // Equatorial Guinea
		NATION_CODE_ERI: "ER", // Eritrea
		NATION_CODE_SWZ: "SZ", // Eswatini
		NATION_CODE_FLK: "FK", // Falkland Islands (Malvinas)
		NATION_CODE_FRO: "FO", // Faroe Islands
		NATION_CODE_FJI: "FJ", // Fiji
		NATION_CODE_GUF: "GF", // French Guiana
		NATION_CODE_PYF: "PF", // French Polynesia
		NATION_CODE_ATF: "TF", // French Southern Territories
		NATION_CODE_GAB: "GA", // Gabon
		NATION_CODE_GMB: "GM", // Gambia
		NATION_CODE_GEO: "GE", // Georgia
		NATION_CODE_GHA: "GH", // Ghana
		NATION_CODE_GIB: "GI", // Gibraltar
		NATION_CODE_GRL: "GL", // Greenland
		NATION_CODE_GRD: "GD", // Grenada
		NATION_CODE_GLP: "GP", // Guadeloupe
		NATION_CODE_GUM: "GU", // Guam
		NATION_CODE_GTM: "GT", // Guatemala
		NATION_CODE_GGY: "GG", // Guernsey
		NATION_CODE_GIN: "GN", // Guinea
		NATION_CODE_GNB: "GW", // Guinea-Bissau
		NATION_CODE_GUY: "GY", // Guyana
		NATION_CODE_HTI: "HT", // Haiti
		NATION_CODE_HMD: "HM", // Heard Island and McDonald Islands
		NATION_CODE_VAT: "VA", // Holy See
		NATION_CODE_HND: "HN", // Honduras
		NATION_CODE_ISL: "IS", // Iceland
		NATION_CODE_IRN: "IR", // Iran (Islamic Republic of)
		NATION_CODE_IRQ: "IQ", // Iraq
		NATION_CODE_IMN: "IM", // Isle of Man
		NATION_CODE_JAM: "JM", // Jamaica
		NATION_CODE_JEY: "JE", // Jersey
		NATION_CODE_JOR: "JO", // Jordan
		NATION_CODE_KAZ: "KZ", // Kazakhstan
		NATION_CODE_KIR: "KI", // Kiribati
		NATION_CODE_PRK: "KP", // Korea (Democratic People's Republic of)
		NATION_CODE_KWT: "KW", // Kuwait
		NATION_CODE_KGZ: "KG", // Kyrgyzstan
		NATION_CODE_LAO: "LA", // Lao People's Democratic Republic
		NATION_CODE_LBN: "LB", // Lebanon
		NATION_CODE_LSO: "LS", // Lesotho
		NATION_CODE_LBR: "LR", // Liberia
		NATION_CODE_LBY: "LY", // Libya
		NATION_CODE_LIE: "LI", // Liechtenstein
		NATION_CODE_MAC: "MO", // Macao
		NATION_CODE_MDG: "MG", // Madagascar
		NATION_CODE_MWI: "MW", // Malawi
		NATION_CODE_MDV: "MV", // Maldives
		NATION_CODE_MLI: "ML", // Mali
		NATION_CODE_MHL: "MH", // Marshall Islands
		NATION_CODE_MTQ: "MQ", // Martinique
		NATION_CODE_MRT: "MR", // Mauritania
		NATION_CODE_MUS: "MU", // Mauritius
		NATION_CODE_MYT: "YT", // Mayotte
		NATION_CODE_FSM: "FM", // Micronesia (Federated States of)
		NATION_CODE_MDA: "MD", // Moldova, Republic of
		NATION_CODE_MCO: "MC", // Monaco
		NATION_CODE_MNG: "MN", // Mongolia
		NATION_CODE_MNE: "ME", // Montenegro
		NATION_CODE_MSR: "MS", // Montserrat
		NATION_CODE_MOZ: "MZ", // Mozambique
		NATION_CODE_MMR: "MM", // Myanmar
		NATION_CODE_NAM: "NA", // Namibia
		NATION_CODE_NRU: "NR", // Nauru
		NATION_CODE_NPL: "NP", // Nepal
		NATION_CODE_NCL: "NC", // New Caledonia
		NATION_CODE_NIC: "NI", // Nicaragua
		NATION_CODE_NER: "NE", // Niger
		NATION_CODE_NIU: "NU", // Niue
		NATION_CODE_NFK: "NF", // Norfolk Island
		NATION_CODE_MKD: "MK", // North Macedonia
		NATION_CODE_MNP: "MP", // Northern Mariana Islands
		NATION_CODE_OMN: "OM", // Oman
		NATION_CODE_PLW: "PW", // Palau
		NATION_CODE_PSE: "PS", // Palestine, State of
		NATION_CODE_PAN: "PA", // Panama
		NATION_CODE_PNG: "PG", // Papua New Guinea
		NATION_CODE_PRY: "PY", // Paraguay
		NATION_CODE_PCN: "PN", // Pitcairn
		NATION_CODE_PRI: "PR", // Puerto Rico
		NATION_CODE_QAT: "QA", // Qatar
		NATION_CODE_REU: "RE", // Réunion
		NATION_CODE_RWA: "RW", // Rwanda
		NATION_CODE_BLM: "BL", // Saint Barthélemy
		NATION_CODE_SHN: "SH", // Saint Helena, Ascension and Tristan da Cunha
		NATION_CODE_KNA: "KN", // Saint Kitts and Nevis
		NATION_CODE_LCA: "LC", // Saint Lucia
		NATION_CODE_MAF: "MF", // Saint Martin (French part)
		NATION_CODE_SPM: "PM", // Saint Pierre and Miquelon
		NATION_CODE_VCT: "VC", // Saint Vincent and the Grenadines
		NATION_CODE_WSM: "WS", // Samoa
		NATION_CODE_SMR: "SM", // San Marino
		NATION_CODE_STP: "ST", // Sao Tome and Principe
		NATION_CODE_SEN: "SN", // Senegal
		NATION_CODE_SRB: "RS", // Serbia
		NATION_CODE_SYC: "SC", // Seychelles
		NATION_CODE_SLE: "SL", // Sierra Leone
		NATION_CODE_SXM: "SX", // Sint Maarten (Dutch part)
		NATION_CODE_SLB: "SB", // Solomon Islands
		NATION_CODE_SOM: "SO", // Somalia
		NATION_CODE_SGS: "GS", // South Georgia and the South Sandwich Islands
		NATION_CODE_SSD: "SS", // South Sudan
		NATION_CODE_LKA: "LK", // Sri Lanka
		NATION_CODE_SDN: "SD", // Sudan
		NATION_CODE_SUR: "SR", // Suriname
		NATION_CODE_SJM: "SJ", // Svalbard and Jan Mayen
		NATION_CODE_SYR: "SY", // Syrian Arab Republic
		NATION_CODE_TJK: "TJ", // Tajikistan
		NATION_CODE_TLS: "TL", // Timor-Leste
		NATION_CODE_TGO: "TG", // Togo
		NATION_CODE_TKL: "TK", // Tokelau
		NATION_CODE_TON: "TO", // Tonga
		NATION_CODE_TTO: "TT", // Trinidad and Tobago
		NATION_CODE_TUN: "TN", // Tunisia
		NATION_CODE_TKM: "TM", // Turkmenistan
		NATION_CODE_TCA: "TC", // Turks and Caicos Islands
		NATION_CODE_TUV: "TV", // Tuvalu
		NATION_CODE_UGA: "UG", // Uganda
		NATION_CODE_UMI: "UM", // United States Minor Outlying Islands
		NATION_CODE_URY: "UY", // Uruguay
		NATION_CODE_VUT: "VU", // Vanuatu
		NATION_CODE_VEN: "VE", // Venezuela (Bolivarian Republic of)
		NATION_CODE_VGB: "VG", // Virgin Islands (British)
		NATION_CODE_VIR: "VI", // Virgin Islands (U.S.)
		NATION_CODE_WLF: "WF", // Wallis and Futuna
		NATION_CODE_ESH: "EH", // Western Sahara
		NATION_CODE_YEM: "YE", // Yemen
		NATION_CODE_ZMB: "ZM", // Zambia
		NATION_CODE_ZWE: "ZW", // Zimbabwe
	}

	return converter[nc]
}
