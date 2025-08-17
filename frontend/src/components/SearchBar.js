import React, { useState } from "react";
import { Input, Radio } from "antd";

import { SEARCH_KEY } from "../constants";

const { Search } = Input;

function SearchBar(props) {
  const [searchType, setSearchType] = useState(SEARCH_KEY.all);
  const [error, setError] = useState("");

  const changeSearchType = (e) => { //点击all之后自动搜索
    const searchType = e.target.value;
    setSearchType(searchType);
    setError("");
    if (searchType === SEARCH_KEY.all) {
      props.handleSearch({ type: searchType, keyword: "" });
    }
  };

  const handleSearch = (value) => {
    //check the corner case 
    if (searchType !== SEARCH_KEY.all && value === "") { //就是选择了keyword和user的情况下，又啥东西都不输入
      setError("Please input your search keyword!");
      return;
    }
    setError("");
    props.handleSearch({ type: searchType, keyword: value });
  };

  return (
    <div className="search-bar">
      <Search
        placeholder="input search text"
        enterButton="Search"
        size="large"
        onSearch={handleSearch}
        disabled={searchType === SEARCH_KEY.all}
      />
      <p className="error-msg">{error}</p>

      <Radio.Group
        onChange={changeSearchType}
        value={searchType}
        className="search-type-group"
      >
        <Radio value={SEARCH_KEY.all}>All</Radio>
        <Radio value={SEARCH_KEY.keyword}>Keyword</Radio>
        <Radio value={SEARCH_KEY.user}>User</Radio>
      </Radio.Group>
    </div>
  );
}

export default SearchBar;
