package mensaapi

type LOCATION int
type FACILITY int

const (
	UNI_ULM    FACILITY = 1 // Universität Ulm
	THU        FACILITY = 2 // Technische Hochschule Ulm
	HS_AALEN   FACILITY = 3 // Hochschule Aalen
	PAED_HS_SG FACILITY = 4 // Pädagogische Hochschule Schwäbisch Gmünd
	HS_GEST_SG FACILITY = 5 // Hochschule für Gestaltung Schwäbisch Gmünd
	HBC_BIB    FACILITY = 6 // HBC. Biberach
	DHS_HDH    FACILITY = 7 // Duale Hochschule Heidenheim
)

const (
	// Universität Ulm
	LOC_MENSA_UNI_ULM        LOCATION = 1  // Mensa in der Uni Ulm
	LOC_CAF_HELMHOLZ_UNI_ULM LOCATION = 16 // Cafeteria Helmholtzstraße
	LOC_CAF_WEST_UNI_ULM     LOCATION = 2  // Cafeteria West
	// LOC_CAF_SOUTH_UNI_ULM          =    // Cafeteria Southside - Eingang Süd
	// LOC_CAF_NORD_UNI_ULM           =    // Cafeteria Nord
	// LOC_AUT_BIB_UNI_ULM            =    // Automatencafeteria in der Bibliothek
	// LOC_CAF_TTU_UNI_ULM            =    // Cafeteria TTU
	// LOC_WESTSIDE_UNI_ULM           =    // WestSide Diner
	// LOC_CAMPUS_UNI_ULM             =    // Campus Diner

	// Technische Hochschule Ulm
	LOC_MENSA_THU          LOCATION = 3  // THU Ulm Mensa
	LOC_ESELSBERG_AUSG_THU LOCATION = 10 // THU Ulm Außenstelle Oberer Eselsberg Essensausgabe
	// LOC_CAF_THU                  =    // THU Ulm Cafeteria
	// LOC_CAF_ESELSBERG_THU        =    // THU Ulm Außenstelle Oberer Eselsberg Cafeteria Kiwi

	// Hochschule Aalen
	LOC_MENSA_HS_AALEN      LOCATION = 7 // Mensa
	LOC_CAF_BURREN_HS_AALEN LOCATION = 5 // Cafeteria in der Hochschule
	// LOC_CAF_HOCH_HS_AALEN         =   // Cafeteria in der Hochschule

	// Pädagogische Hochschule Schwäbisch Gmünd
	LOC_MENSA_PAED_HS_SG LOCATION = 4 // Mensaria PH

	//  Hochschule für Gestaltung Schwäbisch Gmünd
	LOC_HS_GEST_SG LOCATION = 12 // Hochschule für Gestaltung Schwäbisch Gmünd

	// HBC. Biberach
	LOC_HBC_BIB LOCATION = 6 // HBC. Biberach

	// Duale Hochschule Heidenheim
	LOC_DHS_HDH LOCATION = 15 // Duale Hochschule Heidenheim
)
