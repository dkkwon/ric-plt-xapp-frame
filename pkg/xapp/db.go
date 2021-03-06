/*
==================================================================================
  Copyright (c) 2019 AT&T Intellectual Property.
  Copyright (c) 2019 Nokia

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
==================================================================================
*/

package xapp

import (
	rnibentities "gerrit.o-ran-sc.org/r/ric-plt/nodeb-rnib.git/entities"
	rnibreader "gerrit.o-ran-sc.org/r/ric-plt/nodeb-rnib.git/reader"
	sdl "gerrit.o-ran-sc.org/r/ric-plt/sdlgo"
	rnibwriter "gerrit.o-ran-sc.org/r/ric-plt/xapp-frame/pkg/rnib"
	"sync"
	"time"
)

// To be removed later
type SDLStatistics struct{}

var SDLCounterOpts = []CounterOpts{
	{Name: "Stored", Help: "The total number of stored SDL transactions"},
	{Name: "StoreError", Help: "The total number of SDL store errors"},
}

type SDLClient struct {
	db    *sdl.SdlInstance
	stat  map[string]Counter
	mux   sync.Mutex
	ready bool
}

// Alias
type RNIBNodeType = rnibentities.Node_Type
type RNIBGlobalNbId = rnibentities.GlobalNbId
type RNIBNodebInfo = rnibentities.NodebInfo
type RNIBIRNibError = error
type RNIBCells = rnibentities.Cells
type RNIBNbIdentity = rnibentities.NbIdentity
type RNIBCellType = rnibentities.Cell_Type
type RNIBCell = rnibentities.Cell
type RNIBEnb = rnibentities.Enb
type RNIBGnb = rnibentities.Gnb

const RNIBNodeENB = rnibentities.Node_ENB
const RNIBNodeGNB = rnibentities.Node_GNB

type RNIBServedCellInfo = rnibentities.ServedCellInfo
type RNIBNodebInfoEnb = rnibentities.NodebInfo_Enb
type RNIBNodebInfoGnb = rnibentities.NodebInfo_Gnb
type RNIBServedNRCell = rnibentities.ServedNRCell
type RNIBServedNRCellInformation = rnibentities.ServedNRCellInformation
type RNIBNrNeighbourInformation = rnibentities.NrNeighbourInformation

type RNIBClient struct {
	db     *sdl.SdlInstance
	reader rnibreader.RNibReader
	writer rnibwriter.RNibWriter
}

// NewSDLClient returns a new SDLClient.
func NewSDLClient(ns string) *SDLClient {
	if ns == "" {
		ns = "sdl"
	}
	return &SDLClient{
		db:    sdl.NewSdlInstance(ns, sdl.NewDatabase()),
		stat:  Metric.RegisterCounterGroup(SDLCounterOpts, "SDL"),
		ready: false,
	}
}

func (s *SDLClient) TestConnection() {
	// Test DB connection, and wait until ready!
	for {
		if _, err := s.db.GetAll(); err == nil {
			break
		}
		Logger.Warn("Database connection not ready, waiting ...")
		time.Sleep(time.Duration(5 * time.Second))
	}
	s.ready = true
	Logger.Info("Connection to database established!")
}

func (s *SDLClient) IsReady() bool {
	return s.ready
}

func (s *SDLClient) doSet(pairs ...interface{}) (err error) {
	err = s.db.Set(pairs)
	if err != nil {
		s.UpdateStatCounter("StoreError")
	} else {
		s.UpdateStatCounter("Stored")
	}
	return
}

func (s *SDLClient) Store(key string, value interface{}) (err error) {
	return s.doSet(key, value)
}

func (s *SDLClient) MStore(pairs ...interface{}) (err error) {
	return s.doSet(pairs)
}

func (s *SDLClient) Read(key string) (value map[string]interface{}, err error) {
	return s.db.Get([]string{key})
}

func (s *SDLClient) MRead(key []string) (value map[string]interface{}, err error) {
	return s.db.Get(key)
}

func (s *SDLClient) ReadAllKeys(key string) (value []string, err error) {
	return s.db.GetAll()
}

func (s *SDLClient) Subscribe(cb func(string, ...string), channel string) error {
	return s.db.SubscribeChannel(cb, channel)
}

func (s *SDLClient) MSubscribe(cb func(string, ...string), channels ...string) error {
	return s.db.SubscribeChannel(cb, channels...)
}

func (s *SDLClient) StoreAndPublish(channel string, event string, pairs ...interface{}) error {
	return s.db.SetAndPublish([]string{channel, event}, pairs...)
}

func (s *SDLClient) MStoreAndPublish(channelsAndEvents []string, pairs ...interface{}) error {
	return s.db.SetAndPublish(channelsAndEvents, pairs...)
}

func (s *SDLClient) Delete(keys []string) (err error) {
	return s.db.Remove(keys)
}

func (s *SDLClient) Clear() {
	s.db.RemoveAll()
}

func (s *SDLClient) RegisterMetrics() {
	s.stat = Metric.RegisterCounterGroup(SDLCounterOpts, "SDL")
}

func (s *SDLClient) UpdateStatCounter(name string) {
	s.mux.Lock()
	s.stat[name].Inc()
	s.mux.Unlock()
}

func (c *SDLClient) GetStat() (t SDLStatistics) {
	return
}

func NewRNIBClient() *RNIBClient {
	s := sdl.NewSdlInstance("e2Manager", sdl.NewDatabase())
	return &RNIBClient{
		db:     s,
		reader: nil,
		writer: nil,
	}
}

func (r *RNIBClient) Subscribe(cb func(string, ...string), channel string) error {
	return r.db.SubscribeChannel(cb, channel)
}

func (r *RNIBClient) StoreAndPublish(channel string, event string, pairs ...interface{}) error {
	return r.db.SetAndPublish([]string{channel, event}, pairs...)
}

func (r *RNIBClient) GetNodeb(invName string) (*RNIBNodebInfo, RNIBIRNibError) {
	return rnibreader.GetRNibReader(r.db).GetNodeb(invName)
}

func (r *RNIBClient) GetNodebByGlobalNbId(t RNIBNodeType, gid *RNIBGlobalNbId) (*RNIBNodebInfo, RNIBIRNibError) {
	return rnibreader.GetRNibReader(r.db).GetNodebByGlobalNbId(t, gid)
}

func (r *RNIBClient) GetCellList(invName string) (*RNIBCells, RNIBIRNibError) {
	return rnibreader.GetRNibReader(r.db).GetCellList(invName)
}

func (r *RNIBClient) GetListGnbIds() ([]*RNIBNbIdentity, RNIBIRNibError) {
	return rnibreader.GetRNibReader(r.db).GetListGnbIds()
}

func (r *RNIBClient) GetListEnbIds() ([]*RNIBNbIdentity, RNIBIRNibError) {
	return rnibreader.GetRNibReader(r.db).GetListEnbIds()
}

func (r *RNIBClient) GetCountGnbList() (int, RNIBIRNibError) {
	return rnibreader.GetRNibReader(r.db).GetCountGnbList()
}

func (r *RNIBClient) GetCell(invName string, pci uint32) (*RNIBCell, RNIBIRNibError) {
	return rnibreader.GetRNibReader(r.db).GetCell(invName, pci)
}

func (r *RNIBClient) GetCellById(cellType RNIBCellType, cellId string) (*RNIBCell, RNIBIRNibError) {
	return rnibreader.GetRNibReader(r.db).GetCellById(cellType, cellId)
}

func (r *RNIBClient) SaveNodeb(nbIdentity *RNIBNbIdentity, entity *RNIBNodebInfo) RNIBIRNibError {
	return rnibwriter.GetRNibWriter(r.db).SaveNodeb(nbIdentity, entity)
}
