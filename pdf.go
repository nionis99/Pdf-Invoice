package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/jung-kurt/gofpdf"
	"github.com/jung-kurt/gofpdf/contrib/barcode"
)

var (
	Debug                 bool
	subs                  sql.NullString
	addr                  sql.NullString
	town                  sql.NullString
	logo                  sql.NullString
	data_galioja          sql.NullString
	barcode_sql           sql.NullString
	barcode_maxima        sql.NullString
	serija                sql.NullString
	serija_nr             sql.NullString
	kodas                 sql.NullString
	pvm_kodas             sql.NullString
	reg_town              sql.NullString
	reg_addr              sql.NullString
	abonentas_id          sql.NullString
	data_nuo              sql.NullString
	data_iki              sql.NullString
	imokos_kodas          sql.NullString
	data_formuota         sql.NullString
	firma                 sql.NullString
	firma_id              sql.NullString
	firma_adresas         sql.NullString
	firma_kodas           sql.NullString
	firma_pvm_kodas       sql.NullString
	firma_bankas_saskaita sql.NullString
	firma_bankas          sql.NullString
	likutis_tikslinis     sql.NullString
	likutis_skola         sql.NullString
	imoka_tikslinis       sql.NullString
	imoka_skola           sql.NullString
	viso_pries_tiks       sql.NullString
	viso_pries_skol       sql.NullString
	viso_po_tikslinis     sql.NullString
	viso_po_skola         sql.NullString
	moketi                sql.NullString
	moketi_zodziu         sql.NullString
	preke_likutis         sql.NullString
	preke_atideta         sql.NullString
	preke_dalinis         sql.NullString
	///////////////////////////////
	pavadinimas  sql.NullString
	kaina_su_pvm sql.NullString
	nuolaida     sql.NullString
	kiekis       sql.NullString
	kaina_be_pvm sql.NullString
	suma_be_pvm  sql.NullString
	pvm          sql.NullString
	suma_pvm     sql.NullString
	suma_su_pvm  sql.NullString
	avansas      sql.NullString
	skola        sql.NullString
	//////////////////////////////
)
var configFile string
var pwd = os.Getenv("PWD")
var pdf = gofpdf.New("P", "mm", "A4", pwd+"/font/")
var id int

func main() {
	pdf.SetMargins(5, 15, 5)
	pdf.SetAutoPageBreak(true, 20)
	pdf.AddFont("Roboto-Bold", "", "Roboto-Bold.json")
	pdf.AddFont("Roboto-Black", "", "Roboto-Black.json")
	pdf.AddFont("Roboto-Medium", "", "Roboto-Medium.json")
	pdf.AddFont("Roboto-Italic", "", "Roboto-Italic.json")
	pdf.AddFont("Roboto-Thin", "", "Roboto-Thin.json")
	pdf.AddFont("Roboto-Light", "", "Roboto-Light.json")
	pdf.AddFont("Roboto-Regular", "", "Roboto-Regular.json")
	pdf.AliasNbPages("")
	Config() //----- Db config
	pageNumeration()
	header()
	table()
	barcod()
	headerSecond()
	if pdf.PageNo() == 1 || pdf.PageNo() == 0 {
		noNum()
	}
	telReport() //----- Mob. ataskaitu formavimas
	err := pdf.OutputFileAndClose("test.pdf")
	checkErr(err)
}
func pageNumeration() {
	pdf.SetFooterFunc(func() {
		pdf.SetY(-15)
		pdf.SetFont("Arial", "I", 7)
		pdf.SetTextColor(128, 128, 128)
		pdf.CellFormat(0, 8, fmt.Sprintf("Puslapis Nr.:%d/{nb}", pdf.PageNo()),
			"", 0, "R", false, 0, "")
	})
}
func noNum() {
	pdf.SetFooterFunc(func() {
		pdf.SetY(-15)
		pdf.SetFillColor(255, 255, 255)
		pdf.CellFormat(0, 8, "", "", 0, "R", true, 0, "")
	})
}
func header() {
	pdf.AddPage()
	tr := pdf.UnicodeTranslatorFromDescriptor("cp1257")
	Db.QueryRow("SELECT _name,_post_town,_post_address,_firma_logo,_serija,_serija_nr,_reg_town,_reg_address,_kodas,_pvm_kodas,_abonentas_id,_data_nuo::varchar,_data_iki::varchar,_imokos_kodas,_data_formuota::varchar, _firma, _firma_adresas, _firma_kodas, _firma_pvm_kodas, _firma_bankas_saskaita, _firma_bankas, _likutis_tikslinis, _likutis_skola, _imoka_tikslinis, _imoka_skola, _viso_pries_tikslinis, _viso_pries_skola, _viso_po_tikslinis, _viso_po_skola, _moketi, _moketi_zodziu, _barkodas, _barkodas_maxima, _preke_likutis, _preke_atideta, _preke_dalinis, _data_galioja::varchar, _firma_id FROM spls.f_saskaita_2_header($1)", id).Scan(&subs, &town, &addr, &logo, &serija, &serija_nr, &reg_town, &reg_addr, &kodas, &pvm_kodas, &abonentas_id, &data_nuo, &data_iki, &imokos_kodas, &data_formuota, &firma, &firma_adresas, &firma_kodas, &firma_pvm_kodas, &firma_bankas_saskaita, &firma_bankas, &likutis_tikslinis, &likutis_skola, &imoka_tikslinis, &imoka_skola, &viso_pries_tiks, &viso_pries_skol, &viso_po_tikslinis,
		&viso_po_skola, &moketi, &moketi_zodziu, &barcode_sql, &barcode_maxima, &preke_likutis, &preke_atideta, &preke_dalinis, &data_galioja, &firma_id)
	if logo.String == "splius" {
		pdf.Image(pwd+"/images/splius_saskaita_logo.jpg", 21, 35, 36, 11, false, "", 0, "http://www.splius.lt")
		pdf.Image(pwd+"/images/Etaplius_logo_125x125.jpg", 62, 26, 30, 26, false, "", 0, "http://www.etaplius.lt")
	}
	if logo.String == "lansneta" {
		pdf.Image(pwd+"/images/lansneta.jpg", 22, 34, 38, 12, false, "", 0, "http://www.lansneta.lt")
	}
	pdf.SetFont("Roboto-Regular", "", 9)
	pdf.Text(21, 53, tr(subs.String))
	pdf.Text(21, 57, tr(addr.String))
	pdf.Text(21, 61, tr(town.String))
	pdf.SetY(35)
	pdf.SetLeftMargin(114)
	pdf.SetRightMargin(3)
	pdf.MultiCell(0, 5, tr("Klientų patogumui „Splius“ įdiegta „vieno numerio“ sistema, todėl visą informaciją apie „Splius“ teikiamas paslaugas bei pagalbą iškilus problemoms galite gauti paskambinę telefonu 19955. Skambinant iš užsienio prašome rinkti +370-41 55 33 22 SPLIUS visą laiką ieško galimybių kaip tobulinti klientų aptarnavimą, tad laukiame Jūsų pasiūlymų, pageidavimų elektroniniu paštu klientuaptarnavimas@splius.lt"), "0", "L", false)
	pdf.SetMargins(5, 15, 5)
	pdf.SetFont("Roboto-Bold", "", 9)
	pdf.SetY(80)
	pdf.CellFormat(pdf.GetStringWidth(tr("PVM SĄSKAITA-FAKTŪRA")), 5, tr("PVM SĄSKAITA-FAKTŪRA"), "", 0, "L", false, 0, "")
	pdf.SetFont("Roboto-Regular", "", 9)
	pdf.CellFormat(pdf.GetStringWidth(" Serija "), 5, tr(" Serija "), "", 0, "L", false, 0, "")
	pdf.SetFont("Roboto-Bold", "", 9)
	pdf.CellFormat(pdf.GetStringWidth(serija.String+" "), 5, tr(serija.String+" "), "", 0, "L", false, 0, "")
	pdf.SetFont("Roboto-Regular", "", 9)
	pdf.CellFormat(pdf.GetStringWidth("Nr. "), 5, tr("Nr. "), "", 0, "L", false, 0, "")
	pdf.SetFont("Roboto-Bold", "", 9)
	pdf.CellFormat(pdf.GetStringWidth(serija_nr.String), 5, tr(serija_nr.String), "", 1, "L", false, 0, "")
	pdf.SetFont("Roboto-Regular", "", 9)
	pdf.CellFormat(pdf.GetStringWidth(subs.String), 5, tr(subs.String), "", 1, "L", false, 0, "")
	if reg_addr.Valid {
		addr.String = reg_addr.String + ", " + reg_town.String
	} else {
		addr.String = reg_town.String
	}
	pdf.CellFormat(pdf.GetStringWidth(addr.String), 5, tr(addr.String), "", 1, "L", false, 0, "")
	if kodas.Valid {
		pdf.CellFormat(pdf.GetStringWidth(tr("Įmonės kodas ")+kodas.String), 5, tr("Įmonės kodas ")+kodas.String, "", 1, "L", false, 0, "")
	}
	if pvm_kodas.Valid {
		pdf.CellFormat(pdf.GetStringWidth(tr("PVM mokėtojo kodas ")+pvm_kodas.String), 5, tr("PVM mokėtojo kodas ")+pvm_kodas.String, "", 1, "L", false, 0, "")
	}
	pdf.Ln(1)
	pdf.SetFont("Roboto-Bold", "", 9)
	pdf.CellFormat(pdf.GetStringWidth(tr("Mokėtojo kodas ")+abonentas_id.String), 5, tr("Mokėtojo kodas ")+abonentas_id.String, "", 1, "L", false, 0, "")
	pdf.SetFont("Roboto-Regular", "", 9)
	pdf.CellFormat(pdf.GetStringWidth("Suteiktos paslaugos nuo "+data_nuo.String+" iki "+data_iki.String), 5, "Suteiktos paslaugos nuo "+data_nuo.String+" iki "+data_iki.String, "", 0, "L", false, 0, "")
	pdf.SetY(80)
	pdf.SetFont("Roboto-Bold", "", 9)
	pdf.SetX(105)
	pdf.CellFormat(0, 5, tr("ĮMOKOS KODAS: "+imokos_kodas.String), "", 1, "L", false, 0, "")
	pdf.SetX(105)
	pdf.CellFormat(0, 5, data_formuota.String, "", 1, "L", false, 0, "")
	pdf.SetFont("Roboto-Regular", "", 9)
	pdf.SetX(105)
	pdf.CellFormat(0, 5, firma.String, "", 1, "L", false, 0, "")
	pdf.SetX(105)
	pdf.CellFormat(0, 5, tr(firma_adresas.String), "", 1, "L", false, 0, "")
	pdf.SetX(105)
	pdf.CellFormat(0, 5, tr("Įmonės kodas ")+firma_kodas.String, "", 1, "L", false, 0, "")
	pdf.SetX(105)
	pdf.CellFormat(0, 5, tr("PVM mokėtojo kodas ")+firma_pvm_kodas.String, "", 1, "L", false, 0, "")
	pdf.SetX(105)
	pdf.CellFormat(0, 5, tr("A/S "+firma_bankas_saskaita.String+", "+firma_bankas.String), "", 1, "L", false, 0, "")
}

//////////////////////////// Config ///////////////////////////////////////
func Config() {
	flag.StringVar(&configFile, "c", "", "config file path")
	flag.IntVar(&id, "i", 0, "saskaitos id")
	flag.Parse()
	CfgParse(configFile)

	// db init
	err := InitDb()
	checkErr(err)

	if CfgHasKey("debug") {
		Debug = CfgBool("debug")
	}
}

/////////////////////////////// Error check ////////////////////////////////
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

////////////////////////////// Lentele ////////////////////////////////////
func table() {
	tr := pdf.UnicodeTranslatorFromDescriptor("cp1257")
	pdf.SetFont("Roboto-Bold", "", 8)
	pdf.SetY(120)
	pdf.SetX(170)
	pdf.MultiCell(17, 4, tr("Tikslinis\nmokėjimas\navansu"), "1", "C", false)
	pdf.SetY(120)
	pdf.SetX(187)
	pdf.MultiCell(17, 6, tr("Skola / Permoka(-)"), "1", "C", false)
	pdf.SetY(132)
	pdf.CellFormat(165, 4, tr("Likutis mėnesio pradžioje"), "0", 0, "R", false, 0, "")
	pdf.SetFont("Roboto-Regular", "", 7)
	pdf.CellFormat(17, 4, likutis_tikslinis.String, "1", 0, "R", false, 0, "")
	pdf.CellFormat(17, 4, likutis_skola.String, "1", 1, "R", false, 0, "")
	pdf.SetFont("Roboto-Bold", "", 8)
	pdf.CellFormat(165, 4, tr("Gautos įmokos per einamą mėnesį"), "0", 0, "R", false, 0, "")
	pdf.SetFont("Roboto-Regular", "", 7)
	pdf.CellFormat(17, 4, imoka_tikslinis.String, "1", 0, "R", false, 0, "")
	pdf.CellFormat(17, 4, imoka_skola.String, "1", 1, "R", false, 0, "")
	pdf.SetFont("Roboto-Bold", "", 8)
	pdf.CellFormat(165, 4, tr("Viso: "), "0", 0, "R", false, 0, "")
	pdf.SetFont("Roboto-Regular", "", 7)
	pdf.CellFormat(17, 4, viso_pries_tiks.String, "1", 0, "R", false, 0, "")
	pdf.CellFormat(17, 4, viso_pries_skol.String, "1", 1, "R", false, 0, "")
	pdf.SetFont("Roboto-Bold", "", 8)
	////////////////////////// Table header ////////////////////////////
	var x float64 = 0
	pdf.SetY(144)
	tabwid := [10]float64{68, 12, 12, 10, 12, 13, 13, 9, 13, 3}
	pdf.CellFormat(tabwid[0], 12, "Paslauga", "1", 0, "C", false, 0, "")
	x += tabwid[0] + 5
	pdf.SetX(x)
	pdf.MultiCell(tabwid[1], 4, "Kaina\nsu\nPVM", "1", "C", false)
	x += tabwid[1]
	pdf.SetY(144)
	pdf.SetX(x)
	pdf.CellFormat(tabwid[2], 12, "Nuolaida", "1", 0, "C", false, 0, "")
	x += tabwid[2]
	pdf.SetX(x)
	pdf.MultiCell(tabwid[3], 6, "Kiekis\nvnt.", "1", "C", false)
	pdf.SetY(144)
	x += tabwid[3]
	pdf.SetX(x)
	pdf.MultiCell(tabwid[4], 4, "Kaina\nbe\nPVM", "1", "C", false)
	pdf.SetY(144)
	x += tabwid[4]
	pdf.SetX(x)
	pdf.MultiCell(tabwid[5], 4, "Suma\nbe\nPVM", "1", "C", false)
	pdf.SetY(144)
	x += tabwid[5]
	pdf.SetX(x)
	pdf.MultiCell(tabwid[6], 4, "PVM\ntarifas\n%", "1", "C", false)
	pdf.SetY(144)
	x += tabwid[6]
	pdf.SetX(x)
	pdf.CellFormat(tabwid[7], 12, "PVM", "1", 0, "C", false, 0, "")
	pdf.MultiCell(tabwid[8], 4, "Suma\nsu\nPVM", "1", "C", false)
	pdf.SetY(144)
	x += tabwid[7] + tabwid[8]
	pdf.SetX(x)
	pdf.CellFormat(tabwid[9], 4, "", "0", 0, "C", false, 0, "")
	pdf.CellFormat(17, 12, "", "1", 0, "", false, 0, "")
	pdf.CellFormat(17, 12, "", "1", 1, "", false, 0, "")
	rows, err := Db.Query("select * from spls.f_saskaita_2_lines($1)", id)
	checkErr(err)
	///////////////////////////// Table body /////////////////////////
	pdf.SetFont("Roboto-Regular", "", 7)
	for rows.Next() {
		rows.Scan(&pavadinimas, &kaina_su_pvm, &nuolaida, &kiekis, &kaina_be_pvm, &suma_be_pvm, &pvm, &suma_pvm, &suma_su_pvm, &avansas, &skola)
		pdf.CellFormat(tabwid[0], 4, tr(pavadinimas.String), "1", 0, "L", false, 0, "")
		pdf.CellFormat(tabwid[1], 4, kaina_su_pvm.String, "1", 0, "R", false, 0, "")
		pdf.CellFormat(tabwid[2], 4, nuolaida.String, "1", 0, "R", false, 0, "")
		pdf.CellFormat(tabwid[3], 4, kiekis.String, "1", 0, "R", false, 0, "")
		pdf.CellFormat(tabwid[4], 4, kaina_be_pvm.String, "1", 0, "R", false, 0, "")
		pdf.CellFormat(tabwid[5], 4, suma_be_pvm.String, "1", 0, "R", false, 0, "")
		pdf.CellFormat(tabwid[6], 4, pvm.String, "1", 0, "R", false, 0, "")
		pdf.CellFormat(tabwid[7], 4, suma_pvm.String, "1", 0, "R", false, 0, "")
		pdf.CellFormat(tabwid[8], 4, suma_su_pvm.String, "1", 0, "R", false, 0, "")
		pdf.CellFormat(tabwid[9], 4, "", "0", 0, "R", false, 0, "")
		pdf.CellFormat(17, 4, avansas.String, "1", 0, "R", false, 0, "")
		pdf.CellFormat(17, 4, skola.String, "1", 1, "R", false, 0, "")
	}
	//////////////////////////////// Table footer /////////////////////////
	Db.QueryRow("select _nuolaida, _suma_be_pvm, _suma_pvm, _suma_su_pvm, _avansas, _skola from spls.f_saskaita_2_lines_sum($1)", id).Scan(&nuolaida, &suma_be_pvm, &suma_pvm, &suma_su_pvm, &avansas, &skola)
	pdf.CellFormat(tabwid[0], 4, "Viso:", "1", 0, "R", false, 0, "")
	pdf.CellFormat(tabwid[1], 4, "", "1", 0, "R", false, 0, "")
	pdf.CellFormat(tabwid[2], 4, "", "1", 0, "R", false, 0, "")
	pdf.CellFormat(tabwid[3], 4, "", "1", 0, "R", false, 0, "")
	pdf.CellFormat(tabwid[4], 4, "", "1", 0, "R", false, 0, "")
	pdf.CellFormat(tabwid[5], 4, suma_be_pvm.String, "1", 0, "R", false, 0, "")
	pdf.CellFormat(tabwid[6], 4, "", "1", 0, "R", false, 0, "")
	pdf.CellFormat(tabwid[7], 4, suma_pvm.String, "1", 0, "R", false, 0, "")
	pdf.CellFormat(tabwid[8], 4, suma_su_pvm.String, "1", 0, "R", false, 0, "")
	pdf.CellFormat(tabwid[9], 4, "", "0", 0, "R", false, 0, "")
	pdf.CellFormat(17, 4, avansas.String, "1", 0, "R", false, 0, "")
	pdf.CellFormat(17, 4, skola.String, "1", 1, "R", false, 0, "")
	pdf.Ln(2)
	if preke_likutis.Valid {
		pdf.SetFont("Roboto-Bold", "", 8)
		pdf.CellFormat(68, 4, tr("Atidėto mokėjimo už prekę neapmokėtas likutis"), "0", 0, "L", false, 0, "")
		pdf.SetFont("Roboto-Regular", "", 7)
		pdf.CellFormat(24, 4, preke_likutis.String, "1", 0, "R", false, 0, "")
	}
	if preke_atideta.Valid {
		pdf.SetX(5)
		pdf.SetFont("Roboto-Bold", "", 8)
		pdf.CellFormat(165, 4, tr("Atidėtas mokėjimas už prekę"), "0", 0, "R", false, 0, "")
		pdf.SetFont("Roboto-Regular", "", 7)
		pdf.CellFormat(17, 4, "", "1", 0, "R", false, 0, "")
		pdf.CellFormat(17, 4, preke_atideta.String, "1", 1, "R", false, 0, "")

	}
	if preke_dalinis.Valid {
		pdf.SetX(5)
		pdf.SetFont("Roboto-Bold", "", 8)
		pdf.CellFormat(165, 4, tr("Dalinis mokėjimas už prekę"), "0", 0, "R", false, 0, "")
		pdf.SetFont("Roboto-Regular", "", 7)
		pdf.CellFormat(17, 4, "", "1", 0, "R", false, 0, "")
		pdf.CellFormat(17, 4, preke_dalinis.String, "1", 1, "R", false, 0, "")
	}
	pdf.SetX(5)
	pdf.SetFont("Roboto-Bold", "", 8)
	pdf.CellFormat(165, 4, "Viso:", "0", 0, "R", false, 0, "")
	pdf.SetFont("Roboto-Regular", "", 7)
	pdf.CellFormat(17, 4, viso_po_tikslinis.String, "1", 0, "R", false, 0, "")
	pdf.CellFormat(17, 4, viso_po_skola.String, "1", 1, "R", false, 0, "")
	pdf.SetFont("Roboto-Bold", "", 8)
	pdf.CellFormat(165, 4, tr("Mokėtina suma žoždiais - "+moketi_zodziu.String), "0", 1, "R", false, 0, "")
	pdf.CellFormat(165, 4, tr("Mokėtina suma, EUR"), "0", 0, "R", false, 0, "")
	pdf.SetFont("Roboto-Regular", "", 7)
	pdf.CellFormat(34, 4, moketi.String, "1", 1, "R", false, 0, "")
}
func barcod() {
	////////////////////////// Barcode ///////////////////////////////
	if barcode_maxima.Valid || barcode_sql.Valid {
		pdf.SetX(5)
		pdf.CellFormat(0, 16, "", "0", 0, "L", false, 0, "")
		if barcode_maxima.Valid {
			key := barcode.RegisterCode128(pdf, barcode_maxima.String)
			barcode.Barcode(pdf, key, 15, pdf.GetY()+4, 72, 14, false)
			pdf.SetFont("Roboto-Regular", "", 9)
			pdf.SetXY(15, pdf.GetY()+18)
			pdf.CellFormat(72, 3, barcode_maxima.String, "2", 0, "C", false, 0, "")
		}
		if barcode_sql.Valid {
			key := barcode.RegisterCode128(pdf, barcode_sql.String)
			barcode.Barcode(pdf, key, 120, pdf.GetY()-8, 45, 8, false)
			pdf.SetFont("Roboto-Regular", "", 8)
			pdf.SetXY(120, pdf.GetY())
			pdf.CellFormat(45, 3, barcode_sql.String, "2", 0, "C", false, 0, "")
		}
	}
}
func headerSecond() {
	tr := pdf.UnicodeTranslatorFromDescriptor("cp1257")
	pdf.SetY(pdf.GetY() + 10)
	pdf.SetFont("Roboto-Bold", "", 9)
	pdf.CellFormat(pdf.GetStringWidth(tr("PVM SĄSKAITA-FAKTŪRA")), 5, tr("PVM SĄSKAITA-FAKTŪRA"), "", 0, "L", false, 0, "")
	pdf.SetFont("Roboto-Regular", "", 9)
	pdf.CellFormat(pdf.GetStringWidth(" Serija "), 5, tr(" Serija "), "", 0, "L", false, 0, "")
	pdf.SetFont("Roboto-Bold", "", 9)
	pdf.CellFormat(pdf.GetStringWidth(serija.String+" "), 5, tr(serija.String+" "), "", 0, "L", false, 0, "")
	pdf.SetFont("Roboto-Regular", "", 9)
	pdf.CellFormat(pdf.GetStringWidth("Nr. "), 5, tr("Nr. "), "", 0, "L", false, 0, "")
	pdf.SetFont("Roboto-Bold", "", 9)
	pdf.CellFormat(pdf.GetStringWidth(serija_nr.String), 5, tr(serija_nr.String), "", 0, "L", false, 0, "")
	pdf.SetFont("Roboto-Regular", "", 9)
	pdf.CellFormat(pdf.GetStringWidth(tr(" (tęsinys)")), 5, tr(" (tęsinys)"), "", 1, "L", false, 0, "")
	pdf.SetY(pdf.GetY() + 5)
	pdf.SetFont("Roboto-Regular", "", 7)
	pdf.MultiCell(0, 3, tr("Sumokėti už paslaugas galite visuose Lietuvos pašto skyriuose;"), "", "L", false)
	pdf.MultiCell(0, 3, tr("Banko padaliniuose: \"Swedbank\", SEB, DNB, Šiaulių, \"Nordea\", Medicinos, \"Citadele\", \"Paysera\" arba per šių bankų elektroninę bankininystę;"), "", "L", false)
	pdf.MultiCell(0, 3, tr("prekybos centrų \"Maxima\" kasose;"), "", "L", false)
	pdf.MultiCell(0, 3, tr("loterijos \"Perlas\" terminaluose;"), "", "L", false)
	pdf.MultiCell(0, 3, tr("NARVESEN ir Lietuvos spaudos kioskuose;"), "", "L", false)
	pdf.MultiCell(0, 3, tr("per sistemas: www.vienasaskaita.lt, www.manogile.lt bei www.manosplius.lt"), "", "L", false)
	pdf.Ln(3)
	pdf.SetFont("Roboto-Bold", "", 7)
	pdf.MultiCell(0, 3, tr("Įmokas mokant internetu būtina nurodyti: mokėtojo kodą "+abonentas_id.String+", įmokos kodą "+imokos_kodas.String), "", "L", false)
	pdf.Ln(1)
	pdf.SetFont("Roboto-Regular", "", 7)
	pdf.MultiCell(0, 3, tr("Įmonės klientų aptarnavimo skyriuje Jūs galite sumokėti be sąskaitos."), "", "L", false)
	pdf.Ln(1)
	pdf.SetFont("Roboto-Bold", "", 7)
	pdf.MultiCell(0, 3, tr("Informacija, sutrikimų registravimas visą parą tel. 19955."), "", "L", false)
	pdf.Ln(1)
	pdf.SetFont("Roboto-Regular", "", 8)
	pdf.MultiCell(0, 3, tr("SĄSKAITĄ PRAŠOME APMOKĖTI IKI "+data_galioja.String), "", "L", false)
	pdf.Ln(1)
	s, err := strconv.Atoi(viso_pries_skol.String)
	f, err := strconv.Atoi(firma_id.String)
	checkErr(err)
	if s > 0 {
		pdf.SetFont("Roboto-Regular", "", 7)
		pdf.MultiCell(0, 3, tr("Informuojame, kad, esant mėnesio skolai už suteiktas paslaugas, Bendrovė turi teisę sustabdyti paslaugų tiekimą bei reikalauti pagal Sutartį priklausančių mokėjimų, žalos atlygio ir perduoti informaciją apie sutarties turiną tretiesiems asmenims (anstoliams, skolų ieškojimo bendrovėms ir t.t.)."), "", "L", false)
	} else {
		pdf.Ln(3)
		pdf.SetFont("Roboto-Bold", "", 9)
		pdf.MultiCell(0, 5, tr("DĖKOJAME, KAD NAUDOJATĖS MŪSŲ PASLAUGOMIS."), "", "L", false)
	}
	if f == 5 {
		pdf.SetFont("Roboto-Regular", "", 7)
		pdf.MultiCell(0, 5, tr("Sumokėti už paslaugas galite UAB \"Lansneta\" klientų aptarnavimo skyriuje arba pavedimu naudojantis elektronine bankininkyste. Įmokas mokant internetu būtina nurodyti: mokėtojo kodą "+abonentas_id.String+", įmokos kodą "+imokos_kodas.String), "", "L", false)
		pdf.Ln(3)
		pdf.SetFont("Roboto-Bold", "", 7)
		pdf.MultiCell(0, 3, tr("Informacija, sutrikimų registravimas visą parą tel. 8-46 344415."), "", "L", false)
		pdf.Ln(1)
		pdf.SetFont("Roboto-Regular", "", 7)
		pdf.MultiCell(0, 3, tr("Įmonės klientų aptarnavimo skyriuje Jūs galite sumokėti be sąskaitos."), "", "L", false)
		pdf.Ln(1)
		pdf.SetFont("Roboto-Regular", "", 8)
		pdf.MultiCell(0, 3, tr("SĄSKAITAS PRAŠOME APMOKĖTI IKI "+data_galioja.String), "", "L", false)
		pdf.Ln(1)
		if s > 0 {
			pdf.SetFont("Roboto-Regular", "", 7)
			pdf.MultiCell(0, 3, tr("Informuojame, kad, esant mėnesio skolai už suteiktas paslaugas, Bendrovė turi teisę sustabdyti paslaugų tiekimą bei reikalauti pagal Sutartį priklausančių mokėjimų, žalos atlygio ir perduoti informaciją apie sutarties turiną tretiesiems asmenims (anstoliams, skolų ieškojimo bendrovėms ir t.t.)."), "", "L", false)
		} else {
			pdf.Ln(3)
			pdf.SetFont("Roboto-Bold", "", 9)
			pdf.MultiCell(0, 5, tr("DĖKOJAME, KAD NAUDOJATĖS MŪSŲ PASLAUGOMIS."), "", "L", false)
		}
	}
}

/////////////////////////////////////////////////////////////////////////
var (
	numeris       sql.NullString
	planas        sql.NullString
	kreditas      sql.NullString
	viso_suma     sql.NullString
	viso_suma_pvm sql.NullString
	su            float64
	be            float64
	viso          float64
	viso_su       float64
	minus         float64
)

////////////////////////////////////////////////////////////////////////
func telReport() {
	tr := pdf.UnicodeTranslatorFromDescriptor("cp1257")
	rows, err := Db.Query("SELECT _numeris, _planas, _kreditas, _data_nuo::varchar, _data_iki::varchar,	_pvm FROM spls.f_telefonija_ataskaita_header($1,$2)", abonentas_id, data_nuo)
	checkErr(err)
	for rows.Next() {
		rows.Scan(&numeris, &planas, &kreditas, &data_nuo, &data_iki, &pvm)
		viso = 0
		viso_su = 0
		minus = 0
		pdf.AddPage()
		pdf.SetFont("Roboto-Bold", "", 9)
		pdf.CellFormat(pdf.GetStringWidth(tr("PVM SĄSKAITA-FAKTŪRA")), 5, tr("PVM SĄSKAITA-FAKTŪRA"), "", 0, "L", false, 0, "")
		pdf.SetFont("Roboto-Regular", "", 9)
		pdf.CellFormat(pdf.GetStringWidth(" Serija "), 5, tr(" Serija "), "", 0, "L", false, 0, "")
		pdf.SetFont("Roboto-Bold", "", 9)
		pdf.CellFormat(pdf.GetStringWidth(serija.String+" "), 5, tr(serija.String+" "), "", 0, "L", false, 0, "")
		pdf.SetFont("Roboto-Regular", "", 9)
		pdf.CellFormat(pdf.GetStringWidth("Nr. "), 5, tr("Nr. "), "", 0, "L", false, 0, "")
		pdf.SetFont("Roboto-Bold", "", 9)
		pdf.CellFormat(pdf.GetStringWidth(serija_nr.String), 5, tr(serija_nr.String), "", 0, "L", false, 0, "")
		pdf.SetFont("Roboto-Regular", "", 9)
		pdf.CellFormat(pdf.GetStringWidth(tr(" (tęsinys)")), 5, tr(" (tęsinys)"), "", 1, "L", false, 0, "")
		pdf.Ln(15)
		pdf.SetFont("Roboto-Bold", "", 12)
		pdf.CellFormat(180, 5, tr("TELEFONIJOS POKALBIŲ ATASKAITA"), "0", 1, "C", false, 0, "")
		pdf.SetFont("Roboto-Regular", "", 9)
		pdf.Ln(40)
		pdf.CellFormat(70, 5, "Naudojimosi laikotarpis", "1", 0, "L", false, 0, "")
		pdf.CellFormat(50, 5, data_nuo.String+" - "+data_iki.String, "1", 1, "C", false, 0, "")
		pdf.CellFormat(70, 5, "Telefono numeris", "1", 0, "L", false, 0, "")
		pdf.SetFont("Roboto-Bold", "", 9)
		pdf.CellFormat(50, 5, numeris.String, "1", 1, "C", false, 0, "")
		pdf.SetFont("Roboto-Regular", "", 9)
		pdf.CellFormat(70, 5, tr("Mokėjimo planas"), "1", 0, "L", false, 0, "")
		pdf.SetFont("Roboto-Bold", "", 9)
		pdf.CellFormat(50, 5, planas.String, "1", 1, "C", false, 0, "")
		pdf.SetFont("Roboto-Regular", "", 9)
		pdf.CellFormat(70, 5, "Kreditas", "1", 0, "L", false, 0, "")
		pdf.CellFormat(50, 5, kreditas.String, "1", 1, "C", false, 0, "")
		pdf.Ln(12)
		pdf.SetFont("Roboto-Bold", "", 9)
		pdf.CellFormat(70, 5, "Mokestis", "1", 0, "L", false, 0, "")
		if data_nuo.String >= "2015-01-01" {
			pdf.CellFormat(40, 5, "Suma su PVM, Eur", "1", 0, "C", false, 0, "")
			pdf.CellFormat(40, 5, "Suma, Eur", "1", 1, "C", false, 0, "")
		} else {
			pdf.CellFormat(40, 5, "Suma su PVM, Lt", "1", 0, "C", false, 0, "")
			pdf.CellFormat(40, 5, "Suma, Lt", "1", 1, "C", false, 0, "")
		}
		rows, err := Db.Query("SELECT * FROM spls.f_telefonija_ataskaita_detail($1,$2,$3)", numeris, data_nuo, data_iki)
		checkErr(err)
		for rows.Next() {
			rows.Scan(&pavadinimas, &suma_su_pvm, &suma_be_pvm)
			pdf.SetFont("Roboto-Regular", "", 9)
			pdf.CellFormat(70, 4, tr(pavadinimas.String), "1", 0, "L", false, 0, "")
			pdf.CellFormat(40, 4, suma_su_pvm.String, "1", 0, "C", false, 0, "")
			pdf.CellFormat(40, 4, suma_be_pvm.String, "1", 1, "C", false, 0, "")
			su, err := strconv.ParseFloat(suma_su_pvm.String, 64)
			be, err := strconv.ParseFloat(suma_be_pvm.String, 64)
			checkErr(err)
			viso = viso + su
			viso_su = viso_su + be
			minus = viso - viso_su
		}
		pdf.SetFont("Roboto-Bold", "", 9)
		pdf.CellFormat(110, 4, tr("Iš viso"), "0", 0, "R", false, 0, "")
		pdf.CellFormat(40, 4, fmt.Sprintf("%.2f", viso_su), "1", 1, "C", false, 0, "")
		pdf.CellFormat(110, 4, "PVM "+pvm.String+"%", "0", 0, "R", false, 0, "")
		pdf.CellFormat(40, 4, fmt.Sprintf("%.2f", minus), "1", 1, "C", false, 0, "")
		pdf.CellFormat(110, 4, "Suma su PVM", "0", 0, "R", false, 0, "")
		pdf.CellFormat(40, 4, fmt.Sprintf("%.2f", viso), "1", 1, "C", false, 0, "")
	}
}
