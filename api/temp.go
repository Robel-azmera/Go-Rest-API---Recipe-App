package api
//
//import (
//"encoding/json"
//"net/http"
//"strconv"
//"time"
//
//"fmt"
//
//"github.com/gorilla/mux"
//"github.com/yuidegm/soccer/api/entity"
//fixture "github.com/yuidegm/soccer/api/fixture"
//"github.com/yuidegm/soccer/api/utils"
//)
//
//type FixtureApiHandler struct {
//fixtureServices fixture.FixturesServices
//}
//
//func NewFixtureApiHandler(fixturesServices fixture.FixturesServices) *FixtureApiHandler {
//return &FixtureApiHandler{fixtureServices: fixturesServices}
//}
//func (fah *FixtureApiHandler) GetFixture(w http.ResponseWriter, r *http.Request) {
//params := mux.Vars(r)
//id, err := strconv.Atoi(params["id"])
//if err != nil {
//w.Header().Set("Content-Type", "application/json")
//http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
//return
//}
//fxtr, errs := fah.fixtureServices.Fixture(uint32(id))
//
//if len(errs) > 0 {
//w.Header().Set("Content-Type", "application/json")
//http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
//return
//}
//
//output, err := json.MarshalIndent(fxtr, "", "\t\t")
//
//if err != nil {
//fmt.Println("error")
//w.Header().Set("Content-Type", "application/json")
//http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
//return
//}
//
//w.Header().Set("Content-Type", "application/json")
//w.Write(output)
//return
//
//}
//
//func (fah *FixtureApiHandler) GetFixtures(w http.ResponseWriter, r *http.Request) {
//
//fixtures, errs := fah.fixtureServices.Fixtures()
//
//if len(errs) > 0 {
//fmt.Println("error1")
//w.Header().Set("Content-Type", "application/json")
//http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
//return
//}
//
//output, err := json.MarshalIndent(fixtures, "", "\t\t")
//
//if err != nil {
//fmt.Println("error2")
//fmt.Println(err)
//w.Header().Set("Content-Type", "application/json")
//http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
//return
//}
//
//w.Header().Set("Content-Type", "application/json")
//w.Write(output)
//return
//
//}
//
//func (fah *FixtureApiHandler) PostFixture(w http.ResponseWriter, r *http.Request) {
//body := utils.BodyParser(r)
//var fxtr entity.Fixture
//var stringFxtr FixtureWithStringTimeStamp
//err := json.Unmarshal(body, &stringFxtr)
//if err != nil {
//utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
//return
//}
//fxtr.RefereeName = stringFxtr.RefereeName
//fxtr.Clubs = stringFxtr.Clubs
//fxtr.StadiumLatitude = stringFxtr.StadiumLatitude
//fxtr.StadiumLongitude = stringFxtr.StadiumLatitude
//
//i, err := strconv.ParseInt(stringFxtr.StartingDate, 10, 64)
//if err != nil {
//utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
//return
//}
//fxtr.StartingDate = time.Unix(i/1000, 0)
//storedFixture, errs := fah.fixtureServices.StoreFixture(&fxtr)
//
//if len(errs) > 0 {
//w.Header().Set("Content-Type", "application/json")
//http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
//return
//
//}
//output, err := json.MarshalIndent(storedFixture, "", "\t\t")
//
//if err != nil {
//fmt.Println(err)
//w.Header().Set("Content-Type", "application/json")
//http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
//return
//}
//
//w.Header().Set("Content-Type", "application/json")
//_, _ = w.Write(output)
//return
//}
//func (fah *FixtureApiHandler) PutFixture(w http.ResponseWriter, r *http.Request) {
//params := mux.Vars(r)
//id, err := strconv.Atoi(params["id"])
//if err != nil {
//fmt.Println("125")
//w.Header().Set("Content-Type", "application/json")
//http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
//return
//}
//
//g_🅴_𝕥, [28.02.21 10:39]
//fxtr, errs := fah.fixtureServices.Fixture(uint32(id))
//var fixtureWithStringTimeStamp FixtureWithStringTimeStamp
//if len(errs) > 0 {
//fmt.Println("134")
//w.Header().Set("Content-Type", "application/json")
//http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
//return
//}
//body := utils.BodyParser(r)
//err = json.Unmarshal(body, &fixtureWithStringTimeStamp)
//if err != nil {
//fmt.Println("153")
//w.Header().Set("Content-Type", "application/json")
//http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
//return
//}
//fxtr.Id = fixtureWithStringTimeStamp.Id
//fxtr.RefereeName = fixtureWithStringTimeStamp.RefereeName
//fxtr.Clubs = fixtureWithStringTimeStamp.Clubs
//fxtr.StadiumLatitude = fixtureWithStringTimeStamp.StadiumLatitude
//fxtr.StadiumLongitude = fixtureWithStringTimeStamp.StadiumLatitude
//
//i, err := strconv.ParseInt(fixtureWithStringTimeStamp.StartingDate, 10, 64)
//if err != nil {
//fmt.Println("162")
//fmt.Println(err)
//utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
//return
//}
//fxtr.StartingDate = time.Unix(i/1000, 0)
//fxtr, errs = fah.fixtureServices.UpdateFixture(fxtr)
//if len(errs) > 0 {
//fmt.Println("169")
//w.Header().Set("Content-Type", "application/json")
//http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
//return
//}
//output, err := json.MarshalIndent(fxtr, "", "\t\t")
//
//if err != nil {
//fmt.Println("169")
//fmt.Println(err)
//w.Header().Set("Content-Type", "application/json")
//http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
//return
//}
//w.Header().Set("Content-Type", "application/json")
//_, _ = w.Write(output)
//return
//}
//func (fah *FixtureApiHandler) DeleteFixture(w http.ResponseWriter, r *http.Request) {
//
//params := mux.Vars(r)
//id, err := strconv.Atoi(params["id"])
//
//if err != nil {
//w.Header().Set("Content-Type", "application/json")
//http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
//return
//}
//
//_, errs := fah.fixtureServices.DeleteFixture(uint32(id))
//
//if len(errs) > 0 {
//w.Header().Set("Content-Type", "application/json")
//http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
//return
//}
//
//w.Header().Set("Content-Type", "application/json")
//w.WriteHeader(http.StatusNoContent)
//return
//}
//
//type FixtureWithStringTimeStamp struct {
//Id               uint32        gorm:"primary_key;auto_increment js" json:"id"
//StartingDate     string        gorm:"type:varchar(255)" json:"starting_date"
//Clubs            []entity.Club gorm:"many2many:fixture_clubs;auto_preload" json:"clubs"
//StadiumLatitude  float64       json:"stadium_latitude"
//StadiumLongitude float64       json:"stadium_longitude"
//RefereeName      string        gorm:"type:varchar(255)" json:"referee_name"
//}