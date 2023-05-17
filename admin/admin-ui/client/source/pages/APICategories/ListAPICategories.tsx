/*
 *  Copyright (c) 2023, WSO2 LLC. (http://www.wso2.org) All Rights Reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

import Alert from '@mui/material/Alert';
import AlertTitle from '@mui/material/AlertTitle';
import Card from '@mui/material/Card';
import Stack from '@mui/material/Stack';
import axios from 'axios';
import Loader from 'components/Loader';
import PaginatedClientSide from 'components/data-table/PaginatedClientSide';
import React, { useEffect, useState } from 'react';
import { FormattedMessage } from 'react-intl';
import AddUpdateAPICategory from './AddUpdateAPICategory';
import DeleteAPICategory from './DeleteAPICategory';

export default function ListAPICategories() {

  const [data, setData] = useState<{ count: number; list: [{ id: string; name: string; description: string; numberOfAPIs: number; }]; } | null>(null);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string>("");

  const fetchData = () => {
    setLoading(true);
    axios('/api/am/admin/api-categories', {
      method: 'GET',
      withCredentials: true,
    }).then((res) => {
      setData(res.data);
    }).catch((err) => {
      setError(err);
    }).finally(() => {
      setLoading(false);
    });
  };

  useEffect(() => {
    fetchData();
  }, []);

  const columns = React.useMemo(
    () => [
      {
        Header: 'Category Name',
        accessor: 'name',
      },
      {
        Header: 'Description',
        accessor: 'description',
      },
      {
        Header: 'Number Of APIs',
        accessor: 'numberOfAPIs',
      },
      {
        Header: 'Actions',
        accessor: 'actions',
        Cell: (e) => {
          return (
            <Stack direction="row" spacing={1}>
              <AddUpdateAPICategory
                id={e.row.original.id}
                nameProp={e.row.original.name}
                descriptionProp={e.row.original.description}
                updateList={fetchData}
              />
              <DeleteAPICategory
                id={e.row.original.id}
                updateList={fetchData}
              />
            </Stack>
          );
        },
      },
    ],
    []
  )
  if (loading) {
    return <Loader />;
  }
  if (error || data === null || data === undefined) {
    return (
      <Alert severity='error'>
        <AlertTitle>Error</AlertTitle>
        There's an error when fetching API Categories — <strong>check it out!</strong>
      </Alert>
    );
  }
  if (data && data.count === 0) {
    return (
      <Card>
        <AddUpdateAPICategory
          id={undefined}
          nameProp={undefined}
          descriptionProp={undefined}
          updateList={fetchData}
        />
        <FormattedMessage
          id='AdminPages.ApiCategories.List.empty.content.apicategories'
          defaultMessage='Add API Category'
        />
      </Card>
    );
  }
  return (
    <>
      <AddUpdateAPICategory
        id={undefined}
        nameProp={undefined}
        descriptionProp={undefined}
        updateList={fetchData}
      />
      <PaginatedClientSide data={data.list} columns={columns} />
    </>
  )
}
