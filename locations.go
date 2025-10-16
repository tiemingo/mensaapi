package mensaapi

type Location int
type Facility int

const (
	FACILITY_UNI_ULM    Facility = 1 // Universität Ulm
	FACILITY_THU        Facility = 2 // Technische Hochschule Ulm
	FACILITY_HS_AALEN   Facility = 3 // Hochschule Aalen
	FACILITY_PAED_HS_SG Facility = 4 // Pädagogische Hochschule Schwäbisch Gmünd
	FACILITY_HS_GEST_SG Facility = 5 // Hochschule für Gestaltung Schwäbisch Gmünd
	FACILITY_HBC_BIB    Facility = 6 // HBC. Biberach
	FACILITY_DHS_HDH    Facility = 7 // Duale Hochschule Heidenheim
)

const (
	// Universität Ulm
	LOCATION_MENSA_UNI_ULM        Location = 1  // Mensa in der Uni Ulm
	LOCATION_CAF_HELMHOLZ_UNI_ULM Location = 16 // Cafeteria Helmholtzstraße
	LOCATION_CAF_WEST_UNI_ULM     Location = 2  // Cafeteria West
	// LOCATION_CAF_SOUTH_UNI_ULM          =    // Cafeteria Southside - Eingang Süd
	// LOCATION_CAF_NORD_UNI_ULM           =    // Cafeteria Nord
	// LOCATION_AUT_BIB_UNI_ULM            =    // Automatencafeteria in der Bibliothek
	// LOCATION_CAF_TTU_UNI_ULM            =    // Cafeteria TTU
	// LOCATION_WESTSIDE_UNI_ULM           =    // WestSide Diner
	// LOCATION_CAMPUS_UNI_ULM             =    // Campus Diner

	// Technische Hochschule Ulm
	LOCATION_MENSA_THU          Location = 3  // THU Ulm Mensa
	LOCATION_ESELSBERG_AUSG_THU Location = 10 // THU Ulm Außenstelle Oberer Eselsberg Essensausgabe
	// LOCATION_CAF_THU                  =    // THU Ulm Cafeteria
	// LOCATION_CAF_ESELSBERG_THU        =    // THU Ulm Außenstelle Oberer Eselsberg Cafeteria Kiwi

	// Hochschule Aalen
	LOCATION_MENSA_HS_AALEN      Location = 7 // Mensa
	LOCATION_CAF_BURREN_HS_AALEN Location = 5 // Cafeteria in der Hochschule
	// LOCATION_CAF_HOCH_HS_AALEN         =   // Cafeteria in der Hochschule

	// Pädagogische Hochschule Schwäbisch Gmünd
	LOCATION_MENSA_PAED_HS_SG Location = 4 // Mensaria PH

	//  Hochschule für Gestaltung Schwäbisch Gmünd
	LOCATION_HS_GEST_SG Location = 12 // Hochschule für Gestaltung Schwäbisch Gmünd

	// HBC. Biberach
	LOCATION_HBC_BIB Location = 6 // HBC. Biberach

	// Duale Hochschule Heidenheim
	LOCATION_DHS_HDH Location = 15 // Duale Hochschule Heidenheim
)
